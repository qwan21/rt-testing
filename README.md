# rt-testing

Для запуска сервера выполните команду `go run cmd/server/main.go` или используйте Makefile.
По умолчанию http-сервер запускается на порту `8080`. Настройки можно изменить в конфигурационном файле `config/config.yaml`.

Для отправки команд на сервер можно использовать любой HTTP-клиент.
Ниже приведены  команды для управления с указанием маршуртов. 
  

* GET     http://endstream-name/v1/product/get              получить продукт 
* POST    http://endstream-name/api/product/find            возвращает предложение по условию

Объект продукта находится в файле `product.json` 

# Пример
Выполним запрос `GET` `http://127.0.0.1:8080/v1/product/get` на сервер для получения продукта
Ответ: 
`{
    "name": "Игровой",
    "components": [
        {
            "name": "Интернет",
            "isMain": true,
            "prices": [
                {
                    "cost": 100,
                    "priceType": "COST",
                    "ruleApplicabilities": [
                        {
                            "codeName": "technology",
                            "operator": "EQ",
                            "value": "adsl"
                        },
                        {
                            "codeName": "internetSpeed",
                            "operator": "EQ",
                            "value": "10"
                        }
                    ]
                },
                {
                    "cost": 150,
                    "priceType": "COST",
                    "ruleApplicabilities": [
                        {
                            "codeName": "technology",
                            "operator": "EQ",
                            "value": "adsl"
                        },
                        {
                            "codeName": "internetSpeed",
                            "operator": "EQ",
                            "value": "15"
                        }
                    ]
                },
                {
                    "cost": 500,
                    "priceType": "COST",
                    "ruleApplicabilities": [
                        {
                            "codeName": "technology",
                            "operator": "EQ",
                            "value": "xpon"
                        },
                        {
                            "codeName": "internetSpeed",
                            "operator": "EQ",
                            "value": "100"
                        }
                    ]
                },
                {
                    "cost": 900,
                    "priceType": "COST",
                    "ruleApplicabilities": [
                        {
                            "codeName": "technology",
                            "operator": "EQ",
                            "value": "xpon"
                        },
                        {
                            "codeName": "internetSpeed",
                            "operator": "EQ",
                            "value": "200"
                        }
                    ]
                },
                {
                    "cost": 200,
                    "priceType": "COST",
                    "ruleApplicabilities": [
                        {
                            "codeName": "technology",
                            "operator": "EQ",
                            "value": "fttb"
                        },
                        {
                            "codeName": "internetSpeed",
                            "operator": "EQ",
                            "value": "30"
                        }
                    ]
                },
                {
                    "cost": 400,
                    "priceType": "COST",
                    "ruleApplicabilities": [
                        {
                            "codeName": "technology",
                            "operator": "EQ",
                            "value": "fttb"
                        },
                        {
                            "codeName": "internetSpeed",
                            "operator": "EQ",
                            "value": "50"
                        }
                    ]
                },
                {
                    "cost": 600,
                    "priceType": "COST",
                    "ruleApplicabilities": [
                        {
                            "codeName": "technology",
                            "operator": "EQ",
                            "value": "fttb"
                        },
                        {
                            "codeName": "internetSpeed",
                            "operator": "EQ",
                            "value": "200"
                        }
                    ]
                },
                {
                    "cost": 10,
                    "priceType": "Discount",
                    "ruleApplicabilities": [
                        {
                            "codeName": "internetSpeed",
                            "operator": "GTE",
                            "value": "50"
                        }
                    ]
                },
                {
                    "cost": 15,
                    "priceType": "Discount",
                    "ruleApplicabilities": [
                        {
                            "codeName": "internetSpeed",
                            "operator": "GTE",
                            "value": "100"
                        }
                    ]
                }
            ]
        },
        {
            "name": "ADSL Модем",
            "prices": [
                {
                    "cost": 300,
                    "priceType": "COST",
                    "ruleApplicabilities": [
                        {
                            "codeName": "technology",
                            "operator": "EQ",
                            "value": "adsl"
                        }
                    ]
                }
            ]
        }
    ]
}`

Выполним запрос `POST` `http://127.0.0.1:8080/v1/product/find` на сервер для получения предложения по продукту в зависимости от условий
Условие: 
`[
  {
    "ruleName": "technology",
    "value": "adsl"
  },
  {
    "ruleName": "internetSpeed",
    "value": "10"
  }
]`

Результат:
`{
    "name": "Игровой",
    "components": [
        {
            "name": "Интернет",
            "isMain": true,
            "prices": [
                {
                    "cost": 100,
                    "priceType": "COST",
                    "ruleApplicabilities": [
                        {
                            "codeName": "internetSpeed",
                            "operator": "EQ",
                            "value": "10"
                        },
                        {
                            "codeName": "technology",
                            "operator": "EQ",
                            "value": "adsl"
                        }
                    ]
                }
            ]
        },
        {
            "name": "ADSL Модем",
            "prices": [
                {
                    "cost": 300,
                    "priceType": "COST",
                    "ruleApplicabilities": [
                        {
                            "codeName": "technology",
                            "operator": "EQ",
                            "value": "adsl"
                        }
                    ]
                }
            ]
        }
    ],
    "totalCost": {
        "cost": 400
    }
}`
