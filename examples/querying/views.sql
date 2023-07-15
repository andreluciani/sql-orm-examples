CREATE VIEW vw_post_titles_images AS
SELECT posts.id AS post_id,
    posts.title AS post_title,
    images.img_url as image_link
FROM posts_images
    INNER JOIN posts ON posts_images.post_id = posts.id
    INNER JOIN images ON posts_images.image_id = images.id
ORDER BY posts.id;