const forums = require('./forums/client');

const client = forums.Client('http://localhost:8080');

// client.createUser("Dmytro", ["Guitar", "SomeTopic"])
// .then((res) => {
//     console.log(res)
// })
// .catch((e) => {
//     console.log(e.message);
// });

// client.createUser("Nick", ["sea", "marriage", "golang"])
// .then((res) => {
//     console.log(res)
// })
// .catch((e) => {
//     console.log(e.message);
// });

client.listForums()
    .then((list) => {
        console.log(list)
    })
    .catch((e) => {
        console.log(e.message);
});

// база данных активна, так что можно тестировать
// доступно 3 форума:
// Odesa, Topic - sea
// Kyiv, Topic - lol
// Lviv, Topic - choco
// Можно добавлять пользователей и проверять