CREATE TABLE IF NOT EXISTS orders (
    id UUID DEFAULT gen_random_uuid(),
    user_id INT NOT NULL,
    total_amount DECIMAL(10,2) NOT NULL,
    order_status VARCHAR(20) DEFAULT 'pending',
    order_date TIMESTAMPTZ DEFAULT now(),

    CONSTRAINT pk_orders PRIMARY KEY (id)
);

-- ORDER ITEMS
CREATE TABLE IF NOT EXISTS order_items (
    id SERIAL,
    order_id UUID NOT NULL,
    product_id INT NOT NULL,
    quantity INT DEFAULT 1,


    CONSTRAINT pk_order_items PRIMARY KEY (id),
    CONSTRAINT fk_order_id
    FOREIGN KEY (order_id)
    REFERENCES orders(id)
    ON DELETE CASCADE,
    CONSTRAINT fk_product_id
    FOREIGN KEY (product_id)
    REFERENCES public.products(id)
    ON DELETE CASCADE
);