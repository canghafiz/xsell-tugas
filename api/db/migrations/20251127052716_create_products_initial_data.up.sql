-- 1. iPhone 15 Pro Max → subcategory: smartphone
INSERT INTO products (listing_user_id, sub_category_id, title, product_slug, description, price, condition, status)
VALUES (2, (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'smartphone'), 'iPhone 15 Pro Max 256GB', 'iphone-15-pro-max-256gb', 'Brand new, titanium finish. Never used, still sealed with full warranty.', 15999000, 'Like New', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'iphone-15-pro-max-256gb'), '/assets/apple_iphone_15_pro_max_natural_titanium_1_1_2.jpg', true, 1),
       ((SELECT product_id FROM products WHERE product_slug = 'iphone-15-pro-max-256gb'), '/assets/apple_iphone_15_pro_max_natural_titanium_2_1_2.jpg', false, 2);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
SELECT p.product_id, cps.category_product_spec_id, v.spec_value
FROM (SELECT product_id FROM products WHERE product_slug = 'iphone-15-pro-max-256gb') p
         CROSS JOIN (VALUES
                         ('Brand', 'Apple'),
                         ('Model', 'iPhone 15 Pro Max'),
                         ('Screen Size (inches)', '6.7'),
                         ('Rear Camera (MP)', '48'),
                         ('Front Camera (MP)', '12'),
                         ('RAM (GB)', '8'),
                         ('Internal Storage (GB)', '256'),
                         ('Battery Capacity (mAh)', '4422'),
                         ('5G Support', 'Yes'),
                         ('Color', 'Titanium Black')
) AS v(spec_name, spec_value)
         JOIN CategoryProductSpecs cps ON cps.spec_name = v.spec_name AND cps.sub_category_id = (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'smartphone');
INSERT INTO location ( product_id, latitude, longitude, address)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'iphone-15-pro-max-256gb'), -5.3919576, 105.28560, 'Kalibalau Kencana, Bandar Lampung, Lampung, Sumatra, 35133, Indonesia');

-- 2. MacBook Air M2 → subcategory: laptop
INSERT INTO products (listing_user_id, sub_category_id, title, product_slug, description, price, condition, status)
VALUES (2, (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'laptop'), 'MacBook Air M2 13-inch', 'macbook-air-m2-13-inch', 'Used for 3 months, like new condition. Comes with charger and sleeve.', 12500000, 'Like New', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'macbook-air-m2-13-inch'), '/assets/macbook-air-m2-13-inch.webp', true, 1),
       ((SELECT product_id FROM products WHERE product_slug = 'macbook-air-m2-13-inch'), '/assets/macbook-air-m2-13-inch_2.webp', false, 2);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
SELECT p.product_id, cps.category_product_spec_id, v.spec_value
FROM (SELECT product_id FROM products WHERE product_slug = 'macbook-air-m2-13-inch') p
         CROSS JOIN (VALUES
                         ('Brand', 'Apple'),
                         ('Model', 'MacBook Air M2'),
                         ('Screen Size (inches)', '13.6'),
                         ('Processor Brand', 'Apple'),
                         ('Processor Model', 'M2'),
                         ('RAM (GB)', '8'),
                         ('SSD Capacity (GB)', '512'),
                         ('Dedicated GPU', 'No'),
                         ('USB-C Ports', '2'),
                         ('Weight (kg)', '1.24')
) AS v(spec_name, spec_value)
         JOIN CategoryProductSpecs cps ON cps.spec_name = v.spec_name AND cps.sub_category_id = (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'laptop');
INSERT INTO location ( product_id, latitude, longitude, address)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'macbook-air-m2-13-inch'), -5.3942934, 105.27771, 'Kalibalau Kencana, Bandar Lampung, Lampung, Sumatra, 35136, Indonesia');

-- 3. Minimalist 2-Seater Sofa → subcategory: living-room
INSERT INTO products (listing_user_id, sub_category_id, title, product_slug, description, price, condition, status)
VALUES (2, (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'living-room'), 'Minimalist 2-Seater Sofa', 'sofa-minimalis-2-seater', 'Soft fabric sofa, gray color. In excellent condition, selling due to relocation.', 1800000, 'Good', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'sofa-minimalis-2-seater'), '/assets/sofa-minimalis-2-seater.png', true, 1);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
SELECT p.product_id, cps.category_product_spec_id, v.spec_value
FROM (SELECT product_id FROM products WHERE product_slug = 'sofa-minimalis-2-seater') p
         CROSS JOIN (VALUES
                         ('Furniture Type', 'Sofa'),
                         ('Frame Material', 'Wood'),
                         ('Width (cm)', '150'),
                         ('Height (cm)', '85'),
                         ('Depth (cm)', '90'),
                         ('Scratches or Wear', 'None')
) AS v(spec_name, spec_value)
         JOIN CategoryProductSpecs cps ON cps.spec_name = v.spec_name AND cps.sub_category_id = (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'living-room');
INSERT INTO location ( product_id, latitude, longitude, address)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'sofa-minimalis-2-seater'), -5.404886004230271, 105.27490160541201, 'Jagabaya III, Bandar Lampung, Lampung, Sumatra, 35126, Indonesia');


-- 4. Nike Air Jordan 1 Sneakers → subcategory: shoes
INSERT INTO products (listing_user_id, sub_category_id, title, product_slug, description, price, condition, status)
VALUES (2, (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'shoes'), 'Nike Air Jordan 1 Sneakers', 'sepatu-nike-air-jordan-1', 'Authentic, size 42. Worn only twice, 95% condition.', 2500000, 'Like New', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'sepatu-nike-air-jordan-1'), '/assets/sepatu-nike-air-jordan-1.webp', true, 1),
       ((SELECT product_id FROM products WHERE product_slug = 'sepatu-nike-air-jordan-1'), '/assets/sepatu-nike-air-jordan-1_2.jpg', false, 2);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
SELECT p.product_id, cps.category_product_spec_id, v.spec_value
FROM (SELECT product_id FROM products WHERE product_slug = 'sepatu-nike-air-jordan-1') p
         CROSS JOIN (VALUES
                         ('Product Type', 'Sneakers'),
                         ('Brand', 'Nike'),
                         ('Size (S/M/L/XL)', '42'),
                         ('Gender', 'Unisex')
) AS v(spec_name, spec_value)
         JOIN CategoryProductSpecs cps ON cps.spec_name = v.spec_name AND cps.sub_category_id = (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'shoes');
INSERT INTO location ( product_id, latitude, longitude, address)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'sepatu-nike-air-jordan-1'), -5.381006443229868, 105.25170101899629, 'Jalan Mayor Sukardi Hamdani, Pelita, Bandar Lampung, Lampung, Sumatra, 35142, Indonesia');

-- 5. Canon EOS R50 → subcategory: camera
INSERT INTO products (listing_user_id, sub_category_id, title, product_slug, description, price, condition, status)
VALUES (2, (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'camera'), 'Canon EOS R50 Mirrorless Camera', 'kamera-canon-eos-r50', 'Latest Canon mirrorless camera. Body only, excellent condition.', 9500000, 'Like New', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'kamera-canon-eos-r50'), '/assets/kamera-canon-eos-r50.jpg', true, 1),
       ((SELECT product_id FROM products WHERE product_slug = 'kamera-canon-eos-r50'), '/assets/kamera-canon-eos-r50_2.jpg', false, 2);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
SELECT p.product_id, cps.category_product_spec_id, v.spec_value
FROM (SELECT product_id FROM products WHERE product_slug = 'kamera-canon-eos-r50') p
         CROSS JOIN (VALUES
                         ('Brand', 'Canon'),
                         ('Model', 'EOS R50'),
                         ('Release Year', '2023')
) AS v(spec_name, spec_value)
         JOIN CategoryProductSpecs cps ON cps.spec_name = v.spec_name AND cps.sub_category_id = (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'camera');
INSERT INTO location ( product_id, latitude, longitude, address)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'kamera-canon-eos-r50'), -5.386737060415824, 105.25992518510353, 'Jalan Kancil, Sidodadi, Bandar Lampung, Lampung, Sumatra, 35132, Indonesia');

-- 6. Solid Teak Study Desk → subcategory: office-furniture
INSERT INTO products (listing_user_id, sub_category_id, title, product_slug, description, price, condition, status)
VALUES (2, (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'office-furniture'), 'Solid Teak Study Desk', 'meja-belajar-kayu-jati', 'Solid wood desk, 120x60 cm. No scratches or damage.', 2200000, 'Good', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'meja-belajar-kayu-jati'), '/assets/meja-belajar-kayu-jati.webp', true, 1);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
SELECT p.product_id, cps.category_product_spec_id, v.spec_value
FROM (SELECT product_id FROM products WHERE product_slug = 'meja-belajar-kayu-jati') p
         CROSS JOIN (VALUES
                         ('Furniture Type', 'Desk'),
                         ('Primary Material', 'Teak'),
                         ('Width (cm)', '120'),
                         ('Depth (cm)', '60')
) AS v(spec_name, spec_value)
         JOIN CategoryProductSpecs cps ON cps.spec_name = v.spec_name AND cps.sub_category_id = (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'office-furniture');
INSERT INTO location ( product_id, latitude, longitude, address)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'meja-belajar-kayu-jati'), -5.3936658, 105.2583455, 'Kedaton, Bandar Lampung, Lampung, Sumatra, 35123, Indonesia');

-- 7. Gucci Marmont Mini Bag → subcategory: bags-wallets
INSERT INTO products (listing_user_id, sub_category_id, title, product_slug, description, price, condition, status)
VALUES (2, (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'bags-wallets'), 'Gucci Marmont Mini Bag', 'tas-gucci-marmont-mini', 'Authentic Gucci, purchased in Paris. Includes dust bag.', 18000000, 'Like New', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'tas-gucci-marmont-mini'), '/assets/tas-gucci-marmont-mini.jpg', true, 1),
       ((SELECT product_id FROM products WHERE product_slug = 'tas-gucci-marmont-mini'), '/assets/tas-gucci-marmont-mini_1.jpg', false, 2);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
SELECT p.product_id, cps.category_product_spec_id, v.spec_value
FROM (SELECT product_id FROM products WHERE product_slug = 'tas-gucci-marmont-mini') p
         CROSS JOIN (VALUES
                         ('Product Type', 'Handbag'),
                         ('Brand', 'Gucci'),
                         ('Material', 'Leather'),
                         ('Color', 'Black'),
                         ('Gender', 'Women')
) AS v(spec_name, spec_value)
         JOIN CategoryProductSpecs cps ON cps.spec_name = v.spec_name AND cps.sub_category_id = (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'bags-wallets');
INSERT INTO location ( product_id, latitude, longitude, address)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'tas-gucci-marmont-mini'), -5.355928317967699, 105.23655548867288, 'Gang Citra, Bandar Lampung, Lampung Selatan, Lampung, Sumatra, 35143, Indonesia');

-- 8. Samsung Galaxy S23 Ultra → subcategory: smartphone
INSERT INTO products (listing_user_id, sub_category_id, title, product_slug, description, price, condition, status)
VALUES (2, (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'smartphone'), 'Samsung Galaxy S23 Ultra', 'samsung-galaxy-s23-ultra', '512GB, green color. Still under official warranty.', 13500000, 'Like New', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'samsung-galaxy-s23-ultra'), '/assets/samsung-galaxy-s23-ultra.jpg', true, 1),
       ((SELECT product_id FROM products WHERE product_slug = 'samsung-galaxy-s23-ultra'), '/assets/samsung-galaxy-s23-ultra_2.png', false, 2);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
SELECT p.product_id, cps.category_product_spec_id, v.spec_value
FROM (SELECT product_id FROM products WHERE product_slug = 'samsung-galaxy-s23-ultra') p
         CROSS JOIN (VALUES
                         ('Brand', 'Samsung'),
                         ('Model', 'Galaxy S23 Ultra'),
                         ('Screen Size (inches)', '6.8'),
                         ('Rear Camera (MP)', '200'),
                         ('Front Camera (MP)', '12'),
                         ('RAM (GB)', '12'),
                         ('Internal Storage (GB)', '512'),
                         ('5G Support', 'Yes'),
                         ('Color', 'Green')
) AS v(spec_name, spec_value)
         JOIN CategoryProductSpecs cps ON cps.spec_name = v.spec_name AND cps.sub_category_id = (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'smartphone');
INSERT INTO location ( product_id, latitude, longitude, address)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'samsung-galaxy-s23-ultra'), -5.378460983276276, 105.25443851653175, 'Jalan Dempo, Pelita, Bandar Lampung, Lampung, Sumatra, 35142, Indonesia');

-- 9. 3-Door Wardrobe → subcategory: bedroom
INSERT INTO products (listing_user_id, sub_category_id, title, product_slug, description, price, condition, status)
VALUES (2, (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'bedroom'), '3-Door Wardrobe', 'lemari-pakaian-3-pintu', 'Wooden wardrobe with mirror. Large size, fits lots of clothes.', 3500000, 'Good', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'lemari-pakaian-3-pintu'), '/assets/lemari-pakaian-3-pintu.jpeg', true, 1);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
SELECT p.product_id, cps.category_product_spec_id, v.spec_value
FROM (SELECT product_id FROM products WHERE product_slug = 'lemari-pakaian-3-pintu') p
         CROSS JOIN (VALUES
                         ('Furniture Type', 'Wardrobe'),
                         ('Primary Material', 'Plywood'),
                         ('Height (cm)', '200'),
                         ('Width (cm)', '150')
) AS v(spec_name, spec_value)
         JOIN CategoryProductSpecs cps ON cps.spec_name = v.spec_name AND cps.sub_category_id = (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'bedroom');
INSERT INTO location ( product_id, latitude, longitude, address)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'lemari-pakaian-3-pintu'), -5.3692013, 105.2547592, 'Ceker Ganjos, Gang Family VI, Sepang Jaya, Bandar Lampung, Lampung, Sumatra, 35141, Indonesia');
-- 10. Rolex Submariner Watch → subcategory: watches-accessories
INSERT INTO products (listing_user_id, sub_category_id, title, product_slug, description, price, condition, status)
VALUES (2, (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'watches-accessories'), 'Rolex Submariner Watch', 'jam-tangan-rolex-submariner', 'Original, full set with box & papers. Price negotiable!', 180000000, 'Like New', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'jam-tangan-rolex-submariner'), '/assets/jam-tangan-rolex-submariner.webp', true, 1),
       ((SELECT product_id FROM products WHERE product_slug = 'jam-tangan-rolex-submariner'), '/assets/jam-tangan-rolex-submariner_2.jpg', false, 2);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
SELECT p.product_id, cps.category_product_spec_id, v.spec_value
FROM (SELECT product_id FROM products WHERE product_slug = 'jam-tangan-rolex-submariner') p
         CROSS JOIN (VALUES
                         ('Product Type', 'Watch'),
                         ('Brand', 'Rolex'),
                         ('Material', 'Stainless Steel'),
                         ('Color', 'Black'),
                         ('Gender', 'Men')
) AS v(spec_name, spec_value)
         JOIN CategoryProductSpecs cps ON cps.spec_name = v.spec_name AND cps.sub_category_id = (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'watches-accessories');
INSERT INTO location ( product_id, latitude, longitude, address)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'jam-tangan-rolex-submariner'), -5.379210560999783, 105.2577764427381, 'Pelita, Bandar Lampung, Lampung, Sumatra, 35142, Indonesia');

-- 11. iPad Pro 12.9-inch M2 → subcategory: tablet
INSERT INTO products (listing_user_id, sub_category_id, title, product_slug, description, price, condition, status)
VALUES (2, (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'tablet'), 'iPad Pro 12.9-inch M2', 'ipad-pro-12-9-inch-m2', '256GB, Wi-Fi only. Still sealed!', 11000000, 'Like New', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'ipad-pro-12-9-inch-m2'), '/assets/ipad-pro-12-9-inch-m2.webp', true, 1),
       ((SELECT product_id FROM products WHERE product_slug = 'ipad-pro-12-9-inch-m2'), '/assets/ipad-pro-12-9-inch-m2_2.webp', false, 2);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
SELECT p.product_id, cps.category_product_spec_id, v.spec_value
FROM (SELECT product_id FROM products WHERE product_slug = 'ipad-pro-12-9-inch-m2') p
         CROSS JOIN (VALUES
                         ('Brand', 'Apple'),
                         ('Model', 'iPad Pro 12.9'),
                         ('Screen Size (inches)', '12.9'),
                         ('Processor Model', 'M2'),
                         ('RAM (GB)', '8'),
                         ('Internal Storage (GB)', '256')
) AS v(spec_name, spec_value)
         JOIN CategoryProductSpecs cps ON cps.spec_name = v.spec_name AND cps.sub_category_id = (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'tablet');
INSERT INTO location ( product_id, latitude, longitude, address)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'ipad-pro-12-9-inch-m2'), -5.380407121242795, 105.25975065174855, 'Jalan Sultan Haji 1, Pelita, Bandar Lampung, Lampung, Sumatra, 35142, Indonesia');

-- 12. Ergonomic Office Chair → subcategory: office-furniture
INSERT INTO products (listing_user_id, sub_category_id, title, product_slug, description, price, condition, status)
VALUES (2, (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'office-furniture'), 'Ergonomic Office Chair', 'kursi-ergonomis-kantor', 'Premium gaming chair, reclining function. Comfortable for 8-hour use.', 2800000, 'Good', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'kursi-ergonomis-kantor'), '/assets/kursi-ergonomis-kantor.webp', true, 1),
       ((SELECT product_id FROM products WHERE product_slug = 'kursi-ergonomis-kantor'), '/assets/kursi-ergonomis-kantor_2.webp', false, 2);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
SELECT p.product_id, cps.category_product_spec_id, v.spec_value
FROM (SELECT product_id FROM products WHERE product_slug = 'kursi-ergonomis-kantor') p
         CROSS JOIN (VALUES
                         ('Furniture Type', 'Chair'),
                         ('Primary Material', 'Metal'),
                         ('Upholstery Material', 'Synthetic Leather'),
                         ('Primary Color', 'Red & Black'),
                         ('Adjustable Height', 'Yes')
) AS v(spec_name, spec_value)
         JOIN CategoryProductSpecs cps ON cps.spec_name = v.spec_name AND cps.sub_category_id = (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'office-furniture');
INSERT INTO location ( product_id, latitude, longitude, address)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'kursi-ergonomis-kantor'), -5.384883, 105.2585654, 'Gang Cempaka, Pelita, Bandar Lampung, Lampung, Sumatra, 35132, Indonesia');

-- 13. Louis Vuitton Men's Leather Wallet → subcategory: bags-wallets
INSERT INTO products (listing_user_id, sub_category_id, title, product_slug, description, price, condition, status)
VALUES (2, (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'bags-wallets'), 'Louis Vuitton Men''s Leather Wallet', 'dompet-kulit-pria-louis-vuitton', 'Authentic, 99% condition. Comes with small pouch.', 6500000, 'Like New', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'dompet-kulit-pria-louis-vuitton'), '/assets/dompet-kulit-pria-louis-vuitton.jpg', true, 1),
       ((SELECT product_id FROM products WHERE product_slug = 'dompet-kulit-pria-louis-vuitton'), '/assets/dompet-kulit-pria-louis-vuitton_2.webp', false, 2);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
SELECT p.product_id, cps.category_product_spec_id, v.spec_value
FROM (SELECT product_id FROM products WHERE product_slug = 'dompet-kulit-pria-louis-vuitton') p
         CROSS JOIN (VALUES
                         ('Product Type', 'Wallet'),
                         ('Brand', 'Louis Vuitton'),
                         ('Material', 'Leather'),
                         ('Color', 'Brown'),
                         ('Gender', 'Men')
) AS v(spec_name, spec_value)
         JOIN CategoryProductSpecs cps ON cps.spec_name = v.spec_name AND cps.sub_category_id = (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'bags-wallets');
INSERT INTO location ( product_id, latitude, longitude, address)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'dompet-kulit-pria-louis-vuitton'), -5.4539772, 105.2590592, 'Masjid As Shonad, Gang Damai II, Teluk Betung, Bandar Lampung, Lampung, Sumatra, 35223, Indonesia');

-- 14. Dell XPS 13 Laptop → subcategory: laptop
INSERT INTO products (listing_user_id, sub_category_id, title, product_slug, description, price, condition, status)
VALUES (2, (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'laptop'), 'Dell XPS 13 Laptop', 'laptop-dell-xps-13', 'Intel Core i7, 16GB RAM, 512GB SSD. Still very fast.', 9000000, 'Good', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'laptop-dell-xps-13'), '/assets/laptop-dell-xps-13.jpg', true, 1),
       ((SELECT product_id FROM products WHERE product_slug = 'laptop-dell-xps-13'), '/assets/laptop-dell-xps-13_2.webp', false, 2);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
SELECT p.product_id, cps.category_product_spec_id, v.spec_value
FROM (SELECT product_id FROM products WHERE product_slug = 'laptop-dell-xps-13') p
         CROSS JOIN (VALUES
                         ('Brand', 'Dell'),
                         ('Model', 'XPS 13'),
                         ('Screen Size (inches)', '13.4'),
                         ('Processor Brand', 'Intel'),
                         ('Processor Model', 'Core i7'),
                         ('RAM (GB)', '16'),
                         ('SSD Capacity (GB)', '512')
) AS v(spec_name, spec_value)
         JOIN CategoryProductSpecs cps ON cps.spec_name = v.spec_name AND cps.sub_category_id = (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'laptop');
INSERT INTO location ( product_id, latitude, longitude, address)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'laptop-dell-xps-13'), -5.456555643385262, 105.25381791887085, 'Pesawahan, Bandar Lampung, Lampung, Sumatra, 35223, Indonesia');

-- 15. Minimalist Bookshelf → subcategory: living-room
INSERT INTO products (listing_user_id, sub_category_id, title, product_slug, description, price, condition, status)
VALUES (2, (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'living-room'), 'Minimalist Bookshelf', 'rak-buku-minimalis', 'Solid wood shelf, 5 tiers. Perfect for bedroom or living room.', 1500000, 'Good', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'rak-buku-minimalis'), '/assets/rak-buku-minimalis.jpeg', true, 1),
       ((SELECT product_id FROM products WHERE product_slug = 'rak-buku-minimalis'), '/assets/rak-buku-minimalis_2.webp', false, 2);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
SELECT p.product_id, cps.category_product_spec_id, v.spec_value
FROM (SELECT product_id FROM products WHERE product_slug = 'rak-buku-minimalis') p
         CROSS JOIN (VALUES
                         ('Furniture Type', 'Bookshelf'),
                         ('Primary Material', 'Solid Wood'),
                         ('Height (cm)', '180'),
                         ('Design Style', 'Minimalist')
) AS v(spec_name, spec_value)
         JOIN CategoryProductSpecs cps ON cps.spec_name = v.spec_name AND cps.sub_category_id = (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'living-room');
INSERT INTO location ( product_id, latitude, longitude, address)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'rak-buku-minimalis'), -5.465056107112045, 105.24330438333226, 'Lampung, Sumatra, 35223, Indonesia');

-- 16. Ray-Ban Aviator Sunglasses → subcategory: watches-accessories
INSERT INTO products (listing_user_id, sub_category_id, title, product_slug, description, price, condition, status)
VALUES (2, (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'watches-accessories'), 'Ray-Ban Original Aviator Sunglasses', 'kacamata-ray-ban-original', 'Aviator model, polarized lenses. Comes with original case.', 1200000, 'Like New', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'kacamata-ray-ban-original'), '/assets/kacamata-ray-ban-original.jpg', true, 1),
       ((SELECT product_id FROM products WHERE product_slug = 'kacamata-ray-ban-original'), '/assets/kacamata-ray-ban-original_2.webp', false, 2);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
SELECT p.product_id, cps.category_product_spec_id, v.spec_value
FROM (SELECT product_id FROM products WHERE product_slug = 'kacamata-ray-ban-original') p
         CROSS JOIN (VALUES
                         ('Product Type', 'Sunglasses'),
                         ('Brand', 'Ray-Ban'),
                         ('Material', 'Metal'),
                         ('Color', 'Gold/Smoke'),
                         ('Gender', 'Unisex')
) AS v(spec_name, spec_value)
         JOIN CategoryProductSpecs cps ON cps.spec_name = v.spec_name AND cps.sub_category_id = (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'watches-accessories');
INSERT INTO location ( product_id, latitude, longitude, address)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'kacamata-ray-ban-original'), -5.3875533, 104.922813, 'Candi Retno, Pringsewu, Lampung, Sumatra, Indonesia');

-- 17. Sony WH-1000XM5 Headphones → subcategory: headphones
INSERT INTO products (listing_user_id, sub_category_id, title, product_slug, description, price, condition, status)
VALUES (2, (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'headphones'), 'Sony WH-1000XM5 Headphones', 'headphone-sony-wh-1000xm5', 'Best-in-class noise cancellation. Still under warranty, excellent condition.', 3800000, 'Like New', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'headphone-sony-wh-1000xm5'), '/assets/headphone-sony-wh-1000xm5.jpg', true, 1),
       ((SELECT product_id FROM products WHERE product_slug = 'headphone-sony-wh-1000xm5'), '/assets/headphone-sony-wh-1000xm5_2.webp', false, 2);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
SELECT p.product_id, cps.category_product_spec_id, v.spec_value
FROM (SELECT product_id FROM products WHERE product_slug = 'headphone-sony-wh-1000xm5') p
         CROSS JOIN (VALUES
                         ('Brand', 'Sony'),
                         ('Model', 'WH-1000XM5'),
                         ('Bluetooth Version', '5.2')
) AS v(spec_name, spec_value)
         JOIN CategoryProductSpecs cps ON cps.spec_name = v.spec_name AND cps.sub_category_id = (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'headphones');
INSERT INTO location ( product_id, latitude, longitude, address)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'headphone-sony-wh-1000xm5'), -5.334038828588326, 105.19796389290117, 'Pemanggilan, Lampung Selatan, Lampung, Sumatra, 35362, Indonesia');

-- 18. Teak Dining Table → subcategory: dining-room
INSERT INTO products (listing_user_id, sub_category_id, title, product_slug, description, price, condition, status)
VALUES (2, (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'dining-room'), 'Teak Dining Table', 'meja-makan-kayu-jati', 'Seats 6 people, includes chairs. Premium quality.', 8500000, 'Good', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'meja-makan-kayu-jati'), '/assets/meja-makan-kayu-jati.webp', true, 1),
       ((SELECT product_id FROM products WHERE product_slug = 'meja-makan-kayu-jati'), '/assets/meja-makan-kayu-jati_2.jpg', false, 2);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
SELECT p.product_id, cps.category_product_spec_id, v.spec_value
FROM (SELECT product_id FROM products WHERE product_slug = 'meja-makan-kayu-jati') p
         CROSS JOIN (VALUES
                         ('Furniture Type', 'Dining Table'),
                         ('Primary Material', 'Teak'),
                         ('Width (cm)', '180'),
                         ('Depth (cm)', '90')
) AS v(spec_name, spec_value)
         JOIN CategoryProductSpecs cps ON cps.spec_name = v.spec_name AND cps.sub_category_id = (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'dining-room');
INSERT INTO location ( product_id, latitude, longitude, address)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'meja-makan-kayu-jati'), -5.340400823983347, 105.20849084990468, 'Hajimena, Lampung Selatan, Lampung, Sumatra, 35362, Indonesia');

-- 19. Adidas Ultraboost Running Shoes → subcategory: shoes
INSERT INTO products (listing_user_id, sub_category_id, title, product_slug, description, price, condition, status)
VALUES (2, (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'shoes'), 'Adidas Ultraboost Running Shoes', 'sepatu-adidas-ultraboost', 'Authentic, size 41. Used for running 5 times, still like new.', 1500000, 'Like New', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'sepatu-adidas-ultraboost'), '/assets/sepatu-adidas-ultraboost.jpeg', true, 1),
       ((SELECT product_id FROM products WHERE product_slug = 'sepatu-adidas-ultraboost'), '/assets/sepatu-adidas-ultraboost_2.webp', false, 2);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
SELECT p.product_id, cps.category_product_spec_id, v.spec_value
FROM (SELECT product_id FROM products WHERE product_slug = 'sepatu-adidas-ultraboost') p
         CROSS JOIN (VALUES
                         ('Product Type', 'Running Shoes'),
                         ('Brand', 'Adidas'),
                         ('Size (S/M/L/XL)', '41'),
                         ('Gender', 'Unisex')
) AS v(spec_name, spec_value)
         JOIN CategoryProductSpecs cps ON cps.spec_name = v.spec_name AND cps.sub_category_id = (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'shoes');
INSERT INTO location ( product_id, latitude, longitude, address)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'sepatu-adidas-ultraboost'), -5.357281363203193, 105.21078770399349, 'Gedong Meneng, Lampung Selatan, Lampung, Sumatra, 35144, Indonesia');

-- 20. DJI Mini 3 Pro Drone → subcategory: camera
INSERT INTO products (listing_user_id, sub_category_id, title, product_slug, description, price, condition, status)
VALUES (2, (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'camera'), 'DJI Mini 3 Pro Drone', 'drone-dji-mini-3-pro', '4K camera, full set. Total flight time: 10 hours.', 7500000, 'Good', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'drone-dji-mini-3-pro'), '/assets/drone-dji-mini-3-pro.webp', true, 1),
       ((SELECT product_id FROM products WHERE product_slug = 'drone-dji-mini-3-pro'), '/assets/drone-dji-mini-3-pro_2.png', false, 2);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
SELECT p.product_id, cps.category_product_spec_id, v.spec_value
FROM (SELECT product_id FROM products WHERE product_slug = 'drone-dji-mini-3-pro') p
         CROSS JOIN (VALUES
                         ('Brand', 'DJI'),
                         ('Model', 'Mini 3 Pro'),
                         ('Rear Camera (MP)', '48')
) AS v(spec_name, spec_value)
         JOIN CategoryProductSpecs cps ON cps.spec_name = v.spec_name AND cps.sub_category_id = (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'camera');
INSERT INTO location ( product_id, latitude, longitude, address)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'drone-dji-mini-3-pro'), -5.369699275815481, 105.21556116710391, 'Kemiling Permai I, Bandar Lampung, Lampung, Sumatra, 35152, Indonesia');

-- 21. Toyota Fortuner VRZ 2023 → subcategory: cars
INSERT INTO products (listing_user_id, sub_category_id, title, product_slug, description, price, condition, status)
VALUES (2, (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'cars'), 'Toyota Fortuner VRZ 2023', 'toyota-fortuner-vrz-2023', 'December 2022 model, privately owned. Regularly serviced at authorized workshop. 18,000 km on odometer.', 580000000, 'Like New', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'toyota-fortuner-vrz-2023'), '/assets/toyota-fortuner-vrz-2023.webp', true, 1),
       ((SELECT product_id FROM products WHERE product_slug = 'toyota-fortuner-vrz-2023'), '/assets/toyota-fortuner-vrz-2023_2.webp', false, 2);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
SELECT p.product_id, cps.category_product_spec_id, v.spec_value
FROM (SELECT product_id FROM products WHERE product_slug = 'toyota-fortuner-vrz-2023') p
         CROSS JOIN (VALUES
                         ('Brand', 'Toyota'),
                         ('Model', 'Fortuner'),
                         ('Year', '2023'),
                         ('Mileage (km)', '18000'),
                         ('Fuel Type', 'Diesel'),
                         ('Transmission Type', 'Automatic'),
                         ('Registration Status', 'Active')
) AS v(spec_name, spec_value)
         JOIN CategoryProductSpecs cps ON cps.spec_name = v.spec_name AND cps.sub_category_id = (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'cars');
INSERT INTO location ( product_id, latitude, longitude, address)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'toyota-fortuner-vrz-2023'), -5.38072467879826, 105.2095588124965, 'Gunung Terang, Bandar Lampung, Lampung, Sumatra, 35153, Indonesia');

-- 22. Honda ADV 160 Motorcycle → subcategory: motorcycles
INSERT INTO products (listing_user_id, sub_category_id, title, product_slug, description, price, condition, status)
VALUES (2, (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'motorcycles'), 'Honda ADV 160 Motorcycle', 'honda-adv-160-motorcycle', 'Used for 3 months, 2,500 km. Complete with official documents. Matte black color.', 32000000, 'Like New', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'honda-adv-160-motorcycle'), '/assets/honda-adv-160-motorcycle_1.webp', true, 1),
       ((SELECT product_id FROM products WHERE product_slug = 'honda-adv-160-motorcycle'), '/assets/honda-adv-160-motorcycle_2.webp', false, 2);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
SELECT p.product_id, cps.category_product_spec_id, v.spec_value
FROM (SELECT product_id FROM products WHERE product_slug = 'honda-adv-160-motorcycle') p
         CROSS JOIN (VALUES
                         ('Brand', 'Honda'),
                         ('Model', 'ADV'),
                         ('Year', '2024'),
                         ('Mileage (km)', '2500'),
                         ('Engine (cc)', '160'),
                         ('Fuel Type', 'Petrol'),
                         ('Transmission Type', 'Automatic'),
                         ('Registration Status', 'Active')
) AS v(spec_name, spec_value)
         JOIN CategoryProductSpecs cps ON cps.spec_name = v.spec_name AND cps.sub_category_id = (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'motorcycles');
INSERT INTO location ( product_id, latitude, longitude, address)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'honda-adv-160-motorcycle'), -5.386582415063443, 105.20609551749433, 'Jalan Asikin 2, Pesawaran, Lampung, Sumatra, 35155, Indonesia');

-- 23. Mitsubishi Xpander Ultimate 2022 → subcategory: cars
INSERT INTO products (listing_user_id, sub_category_id, title, product_slug, description, price, condition, status)
VALUES (2, (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'cars'), 'Mitsubishi Xpander Ultimate 2022', 'mitsubishi-xpander-ultimate-2022', '2022 model, automatic transmission. Full service history, excellent condition.', 225000000, 'Good', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'mitsubishi-xpander-ultimate-2022'), '/assets/mitsubishi-xpander-ultimate-2022.webp', true, 1),
       ((SELECT product_id FROM products WHERE product_slug = 'mitsubishi-xpander-ultimate-2022'), '/assets/mitsubishi-xpander-ultimate-2022_2.webp', false, 2);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
SELECT p.product_id, cps.category_product_spec_id, v.spec_value
FROM (SELECT product_id FROM products WHERE product_slug = 'mitsubishi-xpander-ultimate-2022') p
         CROSS JOIN (VALUES
                         ('Brand', 'Mitsubishi'),
                         ('Model', 'Xpander'),
                         ('Year', '2022'),
                         ('Mileage (km)', '30000'),
                         ('Fuel Type', 'Petrol'),
                         ('Transmission Type', 'Automatic'),
                         ('Registration Status', 'Active')
) AS v(spec_name, spec_value)
         JOIN CategoryProductSpecs cps ON cps.spec_name = v.spec_name AND cps.sub_category_id = (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'cars');
INSERT INTO location ( product_id, latitude, longitude, address)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'mitsubishi-xpander-ultimate-2022'), -5.406156597399348, 105.24708078682474, 'Sawah Lama, Bandar Lampung, Lampung, Sumatra, 35112, Indonesia');

-- 24. Michelin Primacy 4 Car Tire → subcategory: car-parts
INSERT INTO products (listing_user_id, sub_category_id, title, product_slug, description, price, condition, status)
VALUES (2, (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'car-parts'), 'Michelin Primacy 4 Car Tire 205/55 R16', 'michelin-primacy-4-car-tire-205-55-r16', 'Brand new, never used. Genuine Michelin tire with official warranty.', 1250000, 'New', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'michelin-primacy-4-car-tire-205-55-r16'), '/assets/michelin-primacy-4-car-tire-205-55-r16.webp', true, 1);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
SELECT p.product_id, cps.category_product_spec_id, v.spec_value
FROM (SELECT product_id FROM products WHERE product_slug = 'michelin-primacy-4-car-tire-205-55-r16') p
         CROSS JOIN (VALUES
                         ('Brand', 'Michelin'),
                         ('Model', 'Primacy 4'),
                         ('Variant', '205/55 R16')
) AS v(spec_name, spec_value)
         JOIN CategoryProductSpecs cps ON cps.spec_name = v.spec_name AND cps.sub_category_id = (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'car-parts');
INSERT INTO location ( product_id, latitude, longitude, address)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'michelin-primacy-4-car-tire-205-55-r16'), -5.410080919321371, 105.25328390071212, 'Kebonjeruk, Bandar Lampung, Lampung, Sumatra, 35121, Indonesia');

-- 25. KYT RC7 Carbon Full Face Helmet → subcategory: motorcycle-parts
INSERT INTO products (listing_user_id, sub_category_id, title, product_slug, description, price, condition, status)
VALUES (2, (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'motorcycle-parts'), 'KYT RC7 Carbon Full Face Helmet', 'kyt-rc7-carbon-full-face-helmet', 'SNI & DOT certified. Size M. Never used, still sealed.', 1800000, 'New', 'Available');
INSERT INTO productimages (product_id, image_url, is_primary, order_sequence)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'kyt-rc7-carbon-full-face-helmet'), '/assets/kyt-rc7-carbon-full-face-helmet_1.webp', true, 1),
       ((SELECT product_id FROM products WHERE product_slug = 'kyt-rc7-carbon-full-face-helmet'), '/assets/kyt-rc7-carbon-full-face-helmet_2_2.webp', false, 2);
INSERT INTO productspecs (product_id, category_product_spec_id, spec_value)
SELECT p.product_id, cps.category_product_spec_id, v.spec_value
FROM (SELECT product_id FROM products WHERE product_slug = 'kyt-rc7-carbon-full-face-helmet') p
         CROSS JOIN (VALUES
                         ('Brand', 'KYT'),
                         ('Model', 'RC7'),
                         ('Variant', 'Carbon Full Face'),
                         ('Year', '2024'),
                         ('Color', 'Glossy Black')
) AS v(spec_name, spec_value)
         JOIN CategoryProductSpecs cps ON cps.spec_name = v.spec_name AND cps.sub_category_id = (SELECT sub_category_id FROM SubCategories WHERE sub_category_slug = 'motorcycle-parts');
INSERT INTO location ( product_id, latitude, longitude, address)
VALUES ((SELECT product_id FROM products WHERE product_slug = 'kyt-rc7-carbon-full-face-helmet'), -5.4123742, 105.2570327, 'Durian Payung, Tanjung Karang Barat, Lampung, Sumatra, 35121, Indonesia');