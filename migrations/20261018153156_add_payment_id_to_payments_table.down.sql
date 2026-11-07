-- 2. Recreate the sequence
CREATE SEQUENCE payments_history_payment_id_seq;

-- 3. Change payment_id back to INT
ALTER TABLE payments_history
    ALTER COLUMN payment_id TYPE INT USING payment_id::INT;

-- 4. Restore the default tied to the sequence
ALTER TABLE payments_history
    ALTER COLUMN payment_id SET DEFAULT nextval('payments_history_payment_id_seq');

-- 5. Make the sequence track the current max value (important if rows exist)
SELECT setval('payments_history_payment_id_seq', COALESCE((SELECT MAX(payment_id) FROM payments_history), 1));

-- 6. Restore the primary key
ALTER TABLE payments_history
    ADD CONSTRAINT payments_history_pkey PRIMARY KEY (payment_id);

-- 7. Tie the sequence's ownership back to the column (optional but proper SERIAL behavior)
ALTER SEQUENCE payments_history_payment_id_seq OWNED BY payments_history.payment_id;