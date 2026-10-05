CREATE TABLE IF NOT EXISTS payments_history (
    payment_id SERIAL,
    order_id UUID NOT NULL,
    user_id INT NOT NULL,
    reference VARCHAR(100),
    currency VARCHAR(10) NOT NULL,
    payment_method VARCHAR(50) NOT NULL,
    amount DECIMAL(10,2) NOT NULL,
    payment_status VARCHAR(20) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    CONSTRAINT fk_order_id_payments_history
    FOREIGN KEY (order_id)
    REFERENCES public.orders(id),

    CONSTRAINT pk_payments_history PRIMARY KEY (payment_id),
    CONSTRAINT fk_user_id_payments_history 
    FOREIGN KEY (user_id)
    REFERENCES public.users(id)
);