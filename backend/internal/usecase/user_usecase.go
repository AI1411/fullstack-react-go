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
		// メール履歴保存の失敗はエラーとして返さない
	}

	return nil
}

// generateWelcomeEmailBody はウェルカムメールの本文を生成する
func (u *userUseCase) generateWelcomeEmailBody(userName string) string {
	return fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>農業災害支援システムへようこそ</title>
</head>
<body>
    <div style="max-width: 600px; margin: 0 auto; padding: 20px; font-family: Arial, sans-serif;">
        <h1 style="color: #2c5530;">農業災害支援システムへようこそ</h1>

        <p>%s 様</p>

        <p>この度は農業災害支援システムにご登録いただき、ありがとうございます。</p>

        <p>当システムでは以下のサービスをご利用いただけます：</p>
        <ul>
            <li>災害情報の確認</li>
            <li>被害報告の提出</li>
            <li>支援申請の手続き</li>
            <li>通知の受信</li>
        </ul>

        <p>システムの利用方法についてご不明な点がございましたら、お気軽にお問い合わせください。</p>

        <p>今後ともよろしくお願いいたします。</p>

        <hr style="margin: 30px 0;">
        <p style="color: #666; font-size: 14px;">
            農業災害支援システム 運営事務局<br>
            Email: support@agri-disaster.jp<br>
            Web: https://agri-disaster.jp
        </p>
    </div>
</body>
</html>
	`, userName)
}
