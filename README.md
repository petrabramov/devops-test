# Тестовое задание для DevOps Junior

## Описание
Pipeline для приложения на golang. Производит клонирование кода из репозитория, сборку приложения, тестирование и статический анализ кода, загружает docker-образ в docker hub и развертывает приложение в контейнере с форвардингом 80 внутреннего порта на 8000 внешний. Также имеет возможность уведомлений в Slack на стадиях pipeline. При успешном развертывании веб-приложение доступно по адресу http://hostname.domain:8000/

## Для запуска необходимо
- развернуть Jenkins CI
- создать slack workspace и добавить приложение Jenkins CI для мониторинга
- установить базовые расширения + Slack Notification Plugin
- создать pipeline из SCM https://github.com/petrabramov/devops-test
- задать необходимые для работы credentials и переменные в Jenkinsfile

## Credentials и переменные
<table>
    <tr>
        <td>Переменная</td>
        <td>Значение</td>
    </tr>
    <tr>
        <td>gitBranch</td>
        <td>Выбор ветки из git SCM. Например, 'master'</td>
    </tr>
    <tr>
        <td>gitUrl</td>
        <td>URL до репозитория SCM с проектом</td>
    </tr>
    <tr>
        <td>dockerImage</td>
        <td>Наименование собранного docker-image</td>
    </tr>
    <tr>
        <td>hubCredentialID</td>
        <td>Имя связки credentials в Jenkins для авторизации в Docker Hub</td>
    </tr>
    <tr>
        <td>containerName</td>
        <td>Наименование контейнера с приложением в docker</td>
    </tr>
</table>