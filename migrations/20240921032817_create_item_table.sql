-- +goose Up
-- goose postgres "postgres://devpool:123456789@localhost:5432/mydb" up

-- สร้างตาราง items เพื่อจัดเก็บข้อมูลคำร้องขอเบิก
CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    quantity INT NOT NULL,
    status VARCHAR(50) NOT NULL,
    owner_id INT
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
-- ลบตาราง items ก่อนเพราะมี foreign key เชื่อมโยงกับ users
DROP TABLE IF EXISTS items;

-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
