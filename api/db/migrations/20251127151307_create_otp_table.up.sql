CREATE TABLE otp (
                              otp_id SERIAL PRIMARY KEY,
                              email VARCHAR(100),
                              code VARCHAR(10) NOT NULL,
                              purpose VARCHAR(20) NOT NULL DEFAULT 'email_verification',
                              created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                              expire_at TIMESTAMP GENERATED ALWAYS AS (created_at + INTERVAL '5 minutes') STORED,
                              CONSTRAINT chk_verified_otp_contact CHECK (email IS NOT NULL),
                              CONSTRAINT chk_verified_otp_purpose CHECK (purpose IN ('email_verification', 'password_reset'))
);