-- Add user types
INSERT INTO user_types (type)
VALUES ('basic'),
    ('admin');
-- Add users
INSERT INTO users (
        user_type_id,
        name,
        last_name,
        email,
        password,
        salt,
        created_at,
        updated_at
    )
VALUES (
        1,
        'André',
        'Luciani',
        'andre.luciani@email.com',
        '12345',
        'salt',
        NOW(),
        NOW()
    ),
    (
        1,
        'John',
        'Doe',
        'john.doe@email.com',
        '12345',
        'salt',
        NOW(),
        NOW()
    ),
    (
        1,
        'Priscilla',
        'Scott',
        'priscilla.scott@email.com',
        '12345',
        'salt',
        NOW(),
        NOW()
    );
-- Add posts
INSERT INTO posts (title, content, created_at, updated_at)
VALUES (
        'PostgreSQL 101',
        'This is an example post.',
        NOW(),
        NOW()
    ),
    (
        'Bread Recipe',
        'This is an example post.',
        NOW(),
        NOW()
    ),
    (
        'Will AI take over the world?',
        'This is an example post.',
        NOW(),
        NOW()
    ),
    (
        'How to learn a new technology.',
        'This is an example post.',
        NOW(),
        NOW()
    );
-- Add user_post relations
INSERT INTO user_post (user_id, post_id)
VALUES (1, 1),
    (2, 2),
    (3, 3),
    (3, 4);