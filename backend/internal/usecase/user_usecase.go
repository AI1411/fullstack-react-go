//go:generate mockgen -source=user_usecase.go -destination=../../tests/mock/usecase/user_usecase.mock.go
package usecase

import (
	"context"
	"fmt"
	"net/smtp"
	"time"

	"github.com/aws/smithy-go/ptr"

	"github.com/AI1411/fullstack-react-go/internal/domain/model"
	"github.com/AI1411/fullstack-react-go/internal/infra/datastore"
)

type UserUseCase interface {
	ListUsers(ctx context.Context) ([]*model.User, error)
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, id int32) error
}

type userUseCase struct {
	userRepository         datastore.UserRepository
	emailHistoryRepository datastore.EmailHistoryRepository
}

func NewUserUseCase(
	userRepository datastore.UserRepository,
	emailHistoryRepository datastore.EmailHistoryRepository,
) UserUseCase {
	return &userUseCase{
		userRepository:         userRepository,
		emailHistoryRepository: emailHistoryRepository,
	}
}

func (u *userUseCase) ListUsers(ctx context.Context) ([]*model.User, error) {
	users, err := u.userRepository.Find(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userUseCase) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	user, err := u.userRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUseCase) CreateUser(ctx context.Context, user *model.User) error {
	if err := u.userRepository.Create(ctx, user); err != nil {
		return err
	}

	// ユーザー作成後にウェルカムメールを送信
	if err := u.sendWelcomeEmail(ctx, user); err != nil {
		// メール送信エラーはログに記録するが、ユーザー作成は成功として扱う
		fmt.Printf("ウェルカムメール送信に失敗しました: %v\n", err)
	}

	return nil
}

func (u *userUseCase) UpdateUser(ctx context.Context, user *model.User) error {
	if err := u.userRepository.Update(ctx, user); err != nil {
		return err
	}

	return nil
}

func (u *userUseCase) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user, err := u.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUseCase) DeleteUser(ctx context.Context, id int32) error {
	return u.userRepository.Delete(ctx, id)
}

// MailHogの設定
var (
	smtpHost = "mailhog" // docker-composeで指定したサービス名
	smtpPort = 1025      // MailHogのSMTPポート
	fromAddr = "noreply@agri-disaster.jp"
)

// sendWelcomeEmail はウェルカムメールを送信する
func (u *userUseCase) sendWelcomeEmail(ctx context.Context, user *model.User) error {
	subject := "農業災害支援システムへようこそ"
	body := u.generateWelcomeEmailBody(user.Name)

	// SMTPサーバーのアドレス
	smtpServer := fmt.Sprintf("%s:%d", smtpHost, smtpPort)

	// RFC 5322準拠のメールメッセージを作成
	message := fmt.Sprintf(
		"From: %s\r\n"+
			"To: %s\r\n"+
			"Subject: %s\r\n"+
			"MIME-Version: 1.0\r\n"+
			"Content-Type: text/html; charset=UTF-8\r\n"+
			"\r\n"+
			"%s",
		fromAddr,
		user.Email,
		subject,
		body,
	)

	sendEmailStatus := "sent"
	var errorMessage *string

	// MailHogは認証不要なのでnilを渡す
	if err := smtp.SendMail(smtpServer, nil, fromAddr, []string{user.Email}, []byte(message)); err != nil {
		sendEmailStatus = "failed"
		errorMessage = ptr.String(err.Error())
		fmt.Printf("メール送信エラー: %v\n", err)
	} else {
		fmt.Printf("ウェルカムメール送信成功: %s\n", user.Email)
	}

	// メール送信履歴をデータベースに保存
	emailHistory := &model.EmailHistory{
		UserID:       user.ID,
		Email:        user.Email,
		Subject:      subject,
		EmailType:    "welcome",
		Provider:     "mailhog",
		Status:       sendEmailStatus,
		ErrorMessage: errorMessage,
		SentAt:       time.Now(),
	}

	if err := u.emailHistoryRepository.SaveEmailHistory(ctx, emailHistory); err != nil {
		fmt.Printf("メール履歴保存エラー: %v\n", err)
	}

	return nil
}

// generateWelcomeEmailBody はウェルカムメールの本文を生成する
func (u *userUseCase) generateWelcomeEmailBody(userName string) string {
	return fmt.Sprintf(`
<!-- templates/email/verification.html -->
<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>メールアドレス認証</title>
    <style>
        body {
            font-family: 'Hiragino Sans', 'Yu Gothic', sans-serif;
            line-height: 1.6;
            color: #333;
            max-width: 500px;
            margin: 0 auto;
            padding: 20px;
        }
        .container {
            background-color: #ffffff;
            padding: 20px;
            border: 1px solid #ddd;
            border-radius: 4px;
        }
        .button {
            display: inline-block;
            padding: 12px 24px;
            background-color: #007bff;
            color: white;
            text-decoration: none;
            border-radius: 4px;
            margin: 15px 0;
        }
        .info {
            background-color: #f8f9fa;
            padding: 10px;
            border-radius: 4px;
            margin: 15px 0;
            font-size: 14px;
        }
    </style>
</head>
<body>
    <div class="container">
        <h2>メールアドレス認証</h2>

        <p>%sさん</p>

        下記のボタンをクリックしてメールアドレスの認証を完了してください。</p>

        <div style="text-align: center;">
            <a href="{{.VerificationURL}}" class="button">認証を完了する</a>
        </div>

        <div class="info">
            <strong>注意:</strong> このリンクは1時間で期限切れになります。
        </div>

        <hr>
        <p style="font-size: 12px; color: #666;">
        このメールは自動送信です。<br>
        </p>
    </div>
</body>
</html>
	`, userName)
}
