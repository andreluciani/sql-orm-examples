-- Single Table
SELECT COUNT(image_id),
    post_id
FROM posts_images
GROUP BY post_id
ORDER BY count DESC;
-- On Join
SELECT COUNT(users.id) AS "Users",
    cities.name AS "City"
FROM users
    INNER JOIN cities ON users.city_id = cities.id
GROUP BY cities.name
ORDER BY "Users" DESC;