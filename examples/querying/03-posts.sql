DROP TABLE IF EXISTS posts CASCADE;
CREATE TABLE posts (
    id serial PRIMARY KEY,
    user_id integer NOT NULL,
    title varchar NOT NULL,
    content varchar NOT NULL
);
ALTER TABLE posts
ADD CONSTRAINT posts_fk0 FOREIGN KEY (user_id) REFERENCES users(id);
INSERT INTO posts(user_id, title, content)
VALUES (
        ceil(random() * 30),
        '10 Awesome Ways to Photograph Snakes',
        'Laboriosam eos dignissimos quo aliquid. Corporis eum omnis maxime. Temporibus dolor doloremque eum nisi doloremque. Consequatur reiciendis ut cum enim. Ea aliquam ex eveniet corrupti quo cumque.'
    ),
    (
        ceil(random() * 30),
        '7 Pictures of Rihanna That We Would Rather Forget',
        'Ut rem officia totam. Omnis cum fuga officia dolore in. Deserunt et sit maiores sit vitae ex qui iusto officiis. Iusto labore voluptate nihil quo numquam velit rerum error dicta. Dolorem velit consequatur dolores provident aliquid quia at asperiores quas.'
    ),
    (
        ceil(random() * 30),
        'How to Increase Your Income Using Just Your Ankles.',
        'Sit quae ea ex enim et sit corporis tempore. Debitis et quia nisi cumque numquam sit nobis. Fugit praesentium dolores aut. Consequatur ea veritatis atque velit voluptatem doloribus.'
    ),
    (
        ceil(random() * 30),
        '21 Myths About Snakes Debunked',
        'Perferendis nihil accusantium eum. Voluptate quibusdam qui qui recusandae commodi sunt. Quia eos architecto libero libero sunt. Unde ab ad explicabo assumenda ipsum. Dolor qui dolor.'
    ),
    (
        ceil(random() * 30),
        'Introducing programmer - Who Am I And Why Should You Follow Me',
        'Qui animi qui rem illo tenetur consequuntur rerum deserunt. Dolores voluptatem dolore qui voluptatem alias magnam perspiciatis. Nesciunt ut omnis tempore hic magnam.'
    ),
    (
        ceil(random() * 30),
        'Shape Shifter : Fact versus Fiction',
        'Expedita consequatur molestias nam beatae repellendus. Consequuntur aut dolorem reprehenderit recusandae et voluptatem sunt ratione ipsum. Rerum saepe officia quia eaque quasi tenetur. Harum tenetur qui non qui adipisci necessitatibus autem qui voluptatem.'
    ),
    (
        ceil(random() * 30),
        'Can Snakes Dance : An exploration of Memes',
        'Quasi enim sequi aut ipsam ea eveniet. Minus odio molestias officiis commodi aliquid excepturi necessitatibus. Placeat nesciunt cum ipsum ipsum et laudantium et ut ad.'
    ),
    (
        ceil(random() * 30),
        'Snakes Are the New Black',
        'Rerum et et aut animi voluptatem. Nesciunt dolorem dolor iste deleniti ut. Vel magni ut qui nulla hic ipsum. Cumque animi et nostrum ut odit et voluptas. Error nesciunt eos rerum labore voluptatem dignissimos non.'
    ),
    (
        ceil(random() * 30),
        '20 Dress Reviews in Tweet Form',
        'Quae id atque nam suscipit omnis sint. Numquam veniam in velit et. Eum quis quo minima molestias qui sit repudiandae tempora asperiores.'
    ),
    (
        ceil(random() * 30),
        'From Zero to Shape Shifter - Makeover Tips',
        'Earum minima et quia expedita non asperiores atque aut. Ea voluptas fuga qui rem ratione corporis esse. Animi quo consequatur ut nemo autem a. Quis velit libero. Ut voluptatem nemo doloribus vel excepturi sed. Non ea enim autem recusandae enim.'
    ),
    (
        ceil(random() * 30),
        'How to Make Your Own Admirable Dress for less than £5',
        'Vitae molestias consequatur. Sed nesciunt aut quos. Iusto non ut sint beatae eaque earum non. Labore temporibus illum aut.'
    ),
    (
        ceil(random() * 30),
        'Mickey Mouse - 10 Best Moments',
        'Nulla quisquam possimus rerum qui quos voluptates atque suscipit repellat. Labore sed excepturi consequatur est sit aut id sed. Est numquam quis recusandae ipsa odio repellendus. Eum ut tempore.'
    ),
    (
        ceil(random() * 30),
        'How to Attract More Admirable Subscribers',
        'Dignissimos ut possimus. Necessitatibus maiores rerum. Animi blanditiis perspiciatis. Et deserunt quasi ut reprehenderit occaecati et accusantium ullam. Laborum iusto expedita expedita et. Nihil unde et sequi laudantium voluptatem.'
    ),
    (
        ceil(random() * 30),
        'A Day in the Life of programmer',
        'Facere maiores maiores. Non modi quo sit sapiente. Nihil illo quia ex soluta et. Repudiandae incidunt ut excepturi qui sapiente dolores.'
    ),
    (
        ceil(random() * 30),
        'Unboxing My New Shape Shifter Poo',
        'Voluptatibus enim laudantium repellat reprehenderit dolor. Sed in voluptatum earum. Et voluptatem consequatur sequi. Facilis dolorum quae est dolore recusandae sit voluptatem sapiente. Officiis unde dolorum est perferendis eum illo distinctio sequi ullam. Error hic impedit sit modi.'
    ),
    (
        ceil(random() * 30),
        'The Week: Top Stories About Rihanna',
        'Et qui a ea et quo. Voluptate perspiciatis voluptatum rem est eos debitis amet. Praesentium et a dignissimos id est deserunt. Non nobis sint eum sit nisi. Saepe nostrum unde dolores.'
    ),
    (
        ceil(random() * 30),
        '10 Things You Hveve Always Wanted to Know About the Famous Shape Shifter',
        'Repellat adipisci eius rem illum fuga natus sunt cumque rem. Ipsum ratione minus ad velit velit. Delectus doloremque error. Impedit est fugiat itaque dolor. Et harum nesciunt dicta qui est maiores dolor maxime. Dolorem omnis quisquam.'
    ),
    (
        ceil(random() * 30),
        '7 Unmissable YouTube Channels About Thoughts',
        'Consequuntur eligendi doloribus qui ut quia rerum cum. Sit sed molestiae corrupti voluptatem. Quas est qui mollitia harum.'
    ),
    (
        ceil(random() * 30),
        '10 Things Mickey Mouse Can Teach Us About Thoughts',
        'Minus possimus possimus animi labore sapiente. Rerum aut facilis enim dignissimos dolorum. Labore voluptas eos est est fugit. Ut inventore inventore et quam a et eligendi.'
    ),
    (
        ceil(random() * 30),
        'Mistakes That Snakes Make and How to Avoid Them',
        'Deleniti qui omnis voluptas omnis fugit et. Sit dolorum quas magni a molestiae et. Nostrum dicta quod inventore possimus rerum.'
    ),
    (
        ceil(random() * 30),
        '10 Awesome Ways to Photograph Blue Bottles',
        'Rem aut voluptatem debitis placeat quisquam aut. Omnis et delectus placeat aliquid excepturi ut voluptatem. Dolorem quia odio. Provident temporibus aspernatur expedita impedit qui fugiat et quasi eveniet. Sapiente voluptas ut aut non unde. Earum qui architecto.'
    ),
    (
        ceil(random() * 30),
        '7 Pictures of Paul McCartney That We Would Rather Forget',
        'Cum necessitatibus in in vitae. Molestiae inventore maxime qui quod magni autem. Quia eaque harum qui blanditiis facilis ducimus sunt. Qui dolores cumque aut est sed consequatur fugiat aperiam sed.'
    ),
    (
        ceil(random() * 30),
        'How to Increase Your Income Using Just Your Knees.',
        'Repellendus id neque. Accusamus esse voluptate minus. Explicabo quidem minus. Non error rerum temporibus nesciunt. Quas deserunt facere illum. Minima omnis iste cupiditate.'
    ),
    (
        ceil(random() * 30),
        '21 Myths About Blue bottles Debunked',
        'Dolorem laborum deleniti neque totam aut iure dolorum. Eveniet est vero dolor nisi tenetur quia. Esse reprehenderit iste natus explicabo nihil eos numquam ab doloremque.'
    ),
    (
        ceil(random() * 30),
        'Introducing database - Who Am I And Why Should You Follow Me',
        'Recusandae assumenda quo reiciendis tempora est aut laboriosam voluptas veritatis. Quis amet nobis aliquid ab delectus. Est autem incidunt quae qui rem. Est consequatur fuga et illo qui natus.'
    ),
    (
        ceil(random() * 30),
        'Fallen Angel : Fact versus Fiction',
        'Voluptas consequatur occaecati vitae laboriosam modi ut nihil qui impedit. Et mollitia similique ullam. Asperiores optio voluptates ut ipsa. Iste ex fuga porro architecto et molestiae quia enim quis. Quia asperiores numquam perferendis dolor laborum.'
    ),
    (
        ceil(random() * 30),
        'Can Blue Bottles Dance : An exploration of Memes',
        'Quasi laudantium illo sequi deleniti iure voluptate. Sapiente excepturi et mollitia neque. Sint in quia pariatur. Ea omnis et non sunt dicta dolorem perferendis sapiente nam.'
    ),
    (
        ceil(random() * 30),
        'Blue Bottles Are the New Black',
        'Adipisci harum mollitia quia. Vel possimus similique quaerat architecto a illum rerum accusantium. Eveniet voluptas dolorum deleniti repellat minima similique.'
    ),
    (
        ceil(random() * 30),
        '20 Hat Reviews in Tweet Form',
        'Voluptas praesentium adipisci at nam qui fugiat. Voluptates quasi magni nihil quia repudiandae error dolorem. Qui dolorum sint magni quibusdam temporibus et aut. At sit laboriosam perspiciatis.'
    ),
    (
        ceil(random() * 30),
        'From Zero to Fallen Angel - Makeover Tips',
        'Blanditiis labore eaque voluptatibus architecto quia iure exercitationem dolores. Omnis dolores necessitatibus. Fugit dolor ut a non asperiores blanditiis hic rerum. Veritatis numquam inventore voluptatem libero reprehenderit. Blanditiis vel odio nesciunt neque vel harum aut fugiat totam. Molestiae dolorem adipisci commodi corporis modi nostrum.'
    );