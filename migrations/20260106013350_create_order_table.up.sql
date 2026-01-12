CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    total_amount DECIMAL(10,2) NOT NULL,
    order_status VARCHAR(20) DEFAULT 'pending',
    order_date TIMESTAMPTZ DEFAULT now()
);

-- ORDER ITEMS
CREATE TABLE IF NOT EXISTS order_items (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    product_id INT NOT NULL,
    variant_id INT NOT NULL,
    quantity INT DEFAULT 1,
    price DECIMAL(10,2) NOT NULL,

    CONSTRAINT fk_order_id
    FOREIGN KEY (order_id)
    REFERENCES orders(id)
    ON DELETE CASCADE,
    CONSTRAINT fk_product_id
    FOREIGN KEY (product_id)
    REFERENCES public.products(id)
    ON DELETE CASCADE,
    CONSTRAINT fk_product_variant_id
    FOREIGN KEY (variant_id)
    REFERENCES public.product_variants(id)
    ON DELETE CASCADE
);