CREATE TABLE bills (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    loan_id VARCHAR(40) NOT NULL,
    customer_id VARCHAR(40) NOT NULL,
    start_date DATETIME NOT NULL,
    due_date DATETIME NOT NULL,
    paid_at TIMESTAMP,
    amount INT UNSIGNED,
    status ENUM('waiting_for_payment', 'paid') DEFAULT 'waiting_for_payment',
    payment_id VARCHAR(40)
);
