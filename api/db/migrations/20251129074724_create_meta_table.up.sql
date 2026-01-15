DROP TABLE IF EXISTS content_meta_seo;

CREATE TABLE content_meta_seo (
                                  meta_id      SERIAL PRIMARY KEY,
                                  page_key     VARCHAR(100) NOT NULL,
                                  meta_name    TEXT NOT NULL,      -- 'title', 'description', 'keywords', 'google-site-verification'
                                  meta_value   TEXT NOT NULL,
                                  is_active    BOOLEAN DEFAULT true,
                                  updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX idx_content_meta_seo_page_name ON content_meta_seo(page_key, meta_name);