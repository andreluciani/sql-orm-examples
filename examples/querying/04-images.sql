DROP TABLE IF EXISTS images CASCADE;
CREATE TABLE images (
    id serial PRIMARY KEY,
    img_url varchar NOT NULL
);
INSERT INTO images(img_url)
VALUES ('https://onlink.site/YMrl'),
    ('https://onlink.site/iVhX'),
    ('https://onlink.site/yQCF'),
    ('https://onlink.site/93iP'),
    ('https://onlink.site/AuUT');