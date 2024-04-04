# courses-info
Resolução de desafio técnico

## Exercise
API reference: https://docs.teachable.com/reference
Base API URL: https://developers.teachable.com
API key: 7JbSA3ep6XOMV3t8t7QXuXq9HS79Dwnr

As a creator on Teachable’s platform, you’d like to see some information about your published courses.
Using the endpoints provided by the Public API, output the following information for each published
course within your school:
● Course name
● Course heading
● A list of the names and emails of students actively enrolled in the course

### Desired JSON struct

```
"Name": "Golang do zero",
"Heading": "Curso de Golang para iniciantes",
Alunos: [
    {
    	"Name": "Stephanie",
    	"Email": "steph@email.com",
    },
    {
    	"Name": "Bento",
    	"Email": "bento@email.com",
    },
    {
    	"Name": "Aluno1",
    	"Email": "aluno1@email.com",
    }
]
```

