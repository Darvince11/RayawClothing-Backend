ALTER TABLE payments_history
DROP CONSTRAINT pk_payments_history;

ALTER TABLE payments_history
ALTER COLUMN payment_id DROP DEFAULT;

ALTER TABLE payments_history
ALTER COLUMN payment_id TYPE TEXT,
ALTER COLUMN payment_id SET NOT NULL;

DROP SEQUENCE IF EXISTS payments_history_payment_id_seq;