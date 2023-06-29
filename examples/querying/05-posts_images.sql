DROP TABLE IF EXISTS posts_images CASCADE;
CREATE TABLE posts_images (
    id serial PRIMARY KEY,
    post_id integer NOT NULL,
    image_id integer NOT NULL
);
ALTER TABLE posts_images
ADD CONSTRAINT posts_images_fk0 FOREIGN KEY (post_id) REFERENCES posts(id);
ALTER TABLE posts_images
ADD CONSTRAINT posts_images_fk1 FOREIGN KEY (image_id) REFERENCES images(id);
INSERT INTO posts_images(post_id, image_id)
VALUES (ceil(random() * 30), ceil(random() * 5)),
    (ceil(random() * 30), ceil(random() * 5)),
    (ceil(random() * 30), ceil(random() * 5)),
    (ceil(random() * 30), ceil(random() * 5)),
    (ceil(random() * 30), ceil(random() * 5)),
    (ceil(random() * 30), ceil(random() * 5)),
    (ceil(random() * 30), ceil(random() * 5)),
    (ceil(random() * 30), ceil(random() * 5)),
    (ceil(random() * 30), ceil(random() * 5)),
    (ceil(random() * 30), ceil(random() * 5)),
    (ceil(random() * 30), ceil(random() * 5)),
    (ceil(random() * 30), ceil(random() * 5)),
    (ceil(random() * 30), ceil(random() * 5)),
    (ceil(random() * 30), ceil(random() * 5)),
    (ceil(random() * 30), ceil(random() * 5)),
    (ceil(random() * 30), ceil(random() * 5));