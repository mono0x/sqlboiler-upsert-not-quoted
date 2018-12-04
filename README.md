# How to reproduce

```sh
go get -u github.com/volatiletech/sqlboiler
go get -u github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql
docker-compose -f docker-compose.yml up -d
go generate
GO111MODULES=on go test ./models
docker-compose -f docker-compose.yml down # cleanup
```

```
--- FAIL: TestUpsert (0.00s)
    --- FAIL: TestUpsert/Characters (0.00s)
        character_test.go:708: Unable to upsert Character: models: unable to upsert for character: Error 1064: You have an error in your SQL syntax; check the manual that corresponds to your MySQL server version for the right syntax to use near 'character (`id`, `name`) VALUES (?,?) ON DUPLICATE KEY UPDATE `name` = VALUES(`n' at line 1
        character_test.go:716: want one record, got: 0
        character_test.go:725: Unable to upsert Character: models: unable to upsert for character: Error 1064: You have an error in your SQL syntax; check the manual that corresponds to your MySQL server version for the right syntax to use near 'character (`id`, `name`) VALUES (?,?) ON DUPLICATE KEY UPDATE `name` = VALUES(`n' at line 1
        character_test.go:733: want one record, got: 0
FAIL
FAIL    github.com/mono0x/sqlboiler-upsert-not-quoted/models    0.413s
```
