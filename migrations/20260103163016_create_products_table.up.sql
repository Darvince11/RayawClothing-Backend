CREATE TABLE  IF NOT EXISTS products (
    id SERIAL,
    image_url TEXT NOT NULL,
    product_name VARCHAR(100) NOT NULL,
    product_description TEXT,
    price DECIMAL(10,2) NOT NULL,
    category VARCHAR(20) NOT NULL,
    product_status VARCHAR(20) DEFAULT 'available',
    created_at TIMESTAMPTZ DEFAULT now(),

    CONSTRAINT pk_products PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS product_variants (
    id SERIAL,
    product_id INT NOT NULL,
    product_size TEXT[],
    color TEXT[],

    CONSTRAINT pk_product_variants PRIMARY KEY (id),

    CONSTRAINT fk_product_id
    FOREIGN KEY (product_id)
    REFERENCES public.products(id)
    ON DELETE CASCADE
);
