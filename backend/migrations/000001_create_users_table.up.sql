DROP TABLE IF EXISTS users;
CREATE TABLE IF NOT EXISTS users
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    email      VARCHAR(255) NOT NULL UNIQUE,
    password   VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX IF NOT EXISTS idx_users_email ON users (email);
CREATE INDEX IF NOT EXISTS idx_users_deleted_at ON users (deleted_at);

-- ダミーユーザーデータ挿入
INSERT INTO users (name, email, password) VALUES
('田中太郎', 'taro.tanaka@example.com', '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G'),
('佐藤花子', 'hanako.sato@example.com', '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G'),
('鈴木一郎', 'ichiro.suzuki@example.com', '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G'),
('高橋幸子', 'sachiko.takahashi@example.com', '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G'),
('渡辺雄太', 'yuta.watanabe@example.com', '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G'),
('伊藤美咲', 'misaki.ito@example.com', '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G'),
('山本健太', 'kenta.yamamoto@example.com', '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G'),
('中村洋子', 'yoko.nakamura@example.com', '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G'),
('小林直人', 'naoto.kobayashi@example.com', '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G'),
('加藤千尋', 'chihiro.kato@example.com', '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G'),
('松本龍太郎', 'ryutaro.matsumoto@example.com', '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G'),
('井上真希', 'maki.inoue@example.com', '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G'),
('木村大輔', 'daisuke.kimura@example.com', '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G'),
('林優子', 'yuko.hayashi@example.com', '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G'),
('斎藤拓也', 'takuya.saito@example.com', '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G'),
('清水恵子', 'keiko.shimizu@example.com', '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G'),
('山田隆史', 'takashi.yamada@example.com', '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G'),
('中島裕太', 'yuta.nakajima@example.com', '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G'),
('岡田彩香', 'ayaka.okada@example.com', '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G'),
('後藤光希', 'koki.goto@example.com', '$2a$10$1O3JKkfH7UYl1BV7xQV8huqmARZI6L5ewHN.AgCt/WZtFkN4hL2.G');
