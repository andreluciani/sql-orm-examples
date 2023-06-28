DROP TABLE IF EXISTS images CASCADE;
CREATE TABLE images (
    id serial PRIMARY KEY,
    img_url varchar NOT NULL
);
INSERT INTO images(img_url)
VALUES (
        'https://pixabay.com/get/g4bd8033de3047cc6cc15369e57536746091c2866eaafbfcedca933f46582353d2166aa50828dcbf95c751907e7f435433e6d8ad3259a56330a1bb311b12909f391db548ba0aa44795440282b792d53b3_640.jpg'
    ),
    (
        'https://pixabay.com/get/g8998b2fdc49d6e3eb46f94fbe41e2fc77079e0b8826e3987bbb34d4943d4565c99e594d143f896f6f68fe496c513d03f1c044e795ea336808a0b977fb10438cbad3baa2e9ead1d4fb48c9835cb93dd0d_640.jpg'
    ),
    (
        'https://pixabay.com/get/g1f0963492ebe1a9248d47a461d3613f802b3201e6d7d101728cb7b93c7488e7158eee9d9580856c4f34ff96839d4ce3ade6c7d2cdf51eb43f0d613420f586c21155b1154b49fdbd6a6a833c8a22a8de6_640.jpg'
    ),
    (
        'https://pixabay.com/get/g4cc6e60de570fba692395f37e520d97145793376e2715e1d0f07ea224f0098423db81a7ee4a34dc7ce73220f4e72b0364add9c12c4b5078dce4d4e569b029cde65bbfae3cece73ec9c39016ba3b54212_640.jpg'
    ),
    (
        'https://pixabay.com/get/g44364efccbd2f447744de70cf50bf258978c4b759e42904ba3c230a38f1067487092364b3ade55b93af6b2754d84edd93dc3aa200ccd8c41a130d9156edf3ce18d1117308351217360f3056321112d4f_640.jpg'
    ),
    (
        'https://pixabay.com/get/g5aa235611ba4ceef7fefea12c597c77b70e90a3d097f7a403481097b7e5f9458d6aca2db63a4a3a314633f68117e035e185938fa358e64afa8037c93d27ddfc0f8be7e9aa5b918750238c1ef3f9a2250_640.jpg'
    ),
    (
        'https://pixabay.com/get/g1bccc09d335719bd683b464efcdeba7e53f32bab8d2d08aaaf74f69d3be8014ced19e4a5d03d01b27705ebff6165ccc00e4c6fb5b597a3faa2fe9066c296f18fa4213e78ce13da02b3def03b5ebe46a6_640.jpg'
    ),
    (
        'https://pixabay.com/get/g6c9fc03bdd7b936fd19fdd432767f1cdd14b6c8e582e1199e342df2278e610a38313c02fdc870fdcd78d4ff98653cafc7c70fdda67e6328b23957eee7bb509eea421e2cecdfd18ca9acd98000e69c5cb_640.jpg'
    ),
    (
        'https://pixabay.com/get/gd43e6b65911f57c018a415c801399178d8fbcf4c484834b323ad3a2c466d563f598a90e68009e66a81f5f962a602bb13baece41b7b9ef64249b4b86c413a2903d95fd92995b53a336fb1bdead51c7496_640.jpg'
    ),
    (
        'https://pixabay.com/get/g116a5d116472715b10203562496a8f483caad12f4c805bf460a0b14b3e443bc62daaa5af24db1e27960f29cc027d7910c7a3d3345c59e4c5acc97da8f3d51a7570b243edab992a07045de112033ee335_640.jpg'
    );