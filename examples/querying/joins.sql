-- Left Join
SELECT users.name,
    roles.name
FROM users
    LEFT JOIN roles ON users.role_id = roles.id
LIMIT 5;
-- Left Join using AS clause
SELECT u.name,
    r.name AS role
FROM users AS u
    LEFT JOIN roles AS r ON u.role_id = r.id
LIMIT 5;
-- Left Join (excluding)
SELECT p.title,
    p_i.id
FROM posts as p
    LEFT JOIN posts_images as p_i ON p.id = p_i.post_id
WHERE p_i.id IS NULL;
-- Inner Join
SELECT users.name,
    cities.name AS city
FROM users
    INNER JOIN cities ON users.city_id = cities.id
LIMIT 15;
-- Right Join (excluding)
SELECT cities.name AS city,
    users.name
FROM users
    RIGHT JOIN cities ON users.city_id = cities.id
WHERE city IS NULL;
-- Chained Join - Step 1
SELECT posts.id AS "Post ID",
    posts.title AS "Title"
FROM posts
    INNER JOIN posts_images ON posts.id = posts_images.post_id;
-- Chained Join - Step 2
SELECT posts.id AS "Post ID",
    posts.title AS "Title",
    images.img_url as "Images"
FROM posts_images
    INNER JOIN posts ON posts_images.post_id = posts.id
    INNER JOIN images ON posts_images.image_id = images.id
ORDER BY posts.id;