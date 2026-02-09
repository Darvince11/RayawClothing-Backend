CREATE TABLE IF NOT EXISTS payments_history (
    payment_id SERIAL,
    order_id UUID NOT NULL,
    reference VARCHAR(100) NOT NULL,
    payment_method VARCHAR(50) NOT NULL,
    amount DECIMAL(10,2) NOT NULL,
    payment_status VARCHAR(20) NOT NULL,
    payment_date TIMESTAMPTZ DEFAULT now(),
    CONSTRAINT fk_order_id_payments_history
    FOREIGN KEY (order_id)
    REFERENCES public.orders(id),

    CONSTRAINT pk_payments_history PRIMARY KEY (payment_id)
);