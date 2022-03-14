# Change Log

## [v1.5.2](https://github.com/thewizardplusplus/go-exercises-backend/tree/v1.5.2) (2022-02-27)

## [v1.5.1](https://github.com/thewizardplusplus/go-exercises-backend/tree/v1.5.1) (2021-11-17)

## [v1.5](https://github.com/thewizardplusplus/go-exercises-backend/tree/v1.5) (2021-11-17)

## [v1.4](https://github.com/thewizardplusplus/go-exercises-backend/tree/v1.4) (2021-05-10)

## [v1.3](https://github.com/thewizardplusplus/go-exercises-backend/tree/v1.3) (2021-05-08)

## [v1.2](https://github.com/thewizardplusplus/go-exercises-backend/tree/v1.2) (2021-04-02)

## [v1.1](https://github.com/thewizardplusplus/go-exercises-backend/tree/v1.1) (2021-04-01)

## [v1.0.2](https://github.com/thewizardplusplus/go-exercises-backend/tree/v1.0.2) (2021-03-29)

## [v1.0.1](https://github.com/thewizardplusplus/go-exercises-backend/tree/v1.0.1) (2021-03-27)

## [v1.0](https://github.com/thewizardplusplus/go-exercises-backend/tree/v1.0) (2021-03-26)

## [v1.0-beta](https://github.com/thewizardplusplus/go-exercises-backend/tree/v1.0-beta) (2021-03-25)

Beta of the major version. Implement producing solutions via the [RabbitMQ](https://www.rabbitmq.com/) message broker; add the [wait-for-it](https://github.com/vishnubob/wait-for-it) script to the [Docker](https://www.docker.com/) image.

- databases:
  - storing data in the [PostgreSQL](https://www.postgresql.org/) database:
    - add the `storages.CloseDB()` function;
- interaction with queues:
  - using the [RabbitMQ](https://www.rabbitmq.com/) message broker;
  - common properties:
    - automatic declaring of the used queues;
    - passing of a message data in JSON;
  - operations:
    - producing solutions:
      - concurrent producing;
- distributing:
  - [Docker](https://www.docker.com/) image:
    - add the [wait-for-it](https://github.com/vishnubob/wait-for-it) script.

### Features

- RESTful API:
  - models:
    - task model:
      - storing:
        - title;
        - description;
        - boilerplate code;
        - test cases:
          - all test cases are represented by a single string;
      - operations:
        - getting all tasks;
        - getting a single task by an ID;
        - creating;
        - updating by an ID;
        - deleting by an ID;
    - solution model:
      - storing:
        - task ID;
        - code;
      - operations:
        - getting all solutions by a task ID;
        - getting a single solution by an ID;
        - creating;
  - representing:
    - in a JSON:
      - payloads:
        - of requests;
        - of responses;
    - as a plain text:
      - errors;
- server:
  - storing settings in environment variables;
  - supporting graceful shutdown;
  - logging:
    - logging requests;
    - logging errors;
  - panics:
    - recovering on panics;
    - logging of panics;
- databases:
  - storing data in the [PostgreSQL](https://www.postgresql.org/) database;
- interaction with queues:
  - using the [RabbitMQ](https://www.rabbitmq.com/) message broker;
  - common properties:
    - automatic declaring of the used queues;
    - passing of a message data in JSON;
  - operations:
    - producing solutions:
      - concurrent producing;
- distributing:
  - [Docker](https://www.docker.com/) image;
  - [Docker Compose](https://docs.docker.com/compose/) configuration.

## [v1.0-alpha.1](https://github.com/thewizardplusplus/go-exercises-backend/tree/v1.0-alpha.1) (2021-03-24)

Second alpha of the major version. Implement the solution model.

- RESTful API:
  - models:
    - solution model:
      - storing:
        - task ID;
        - code;
      - operations:
        - getting all solutions by a task ID;
        - getting a single solution by an ID;
        - creating.

### Features

- RESTful API:
  - models:
    - task model:
      - storing:
        - title;
        - description;
        - boilerplate code;
        - test cases:
          - all test cases are represented by a single string;
      - operations:
        - getting all tasks;
        - getting a single task by an ID;
        - creating;
        - updating by an ID;
        - deleting by an ID;
    - solution model:
      - storing:
        - task ID;
        - code;
      - operations:
        - getting all solutions by a task ID;
        - getting a single solution by an ID;
        - creating;
  - representing:
    - in a JSON:
      - payloads:
        - of requests;
        - of responses;
    - as a plain text:
      - errors;
- server:
  - storing settings in environment variables;
  - supporting graceful shutdown;
  - logging:
    - logging requests;
    - logging errors;
  - panics:
    - recovering on panics;
    - logging of panics;
- databases:
  - storing data in the [PostgreSQL](https://www.postgresql.org/) database;
- distributing:
  - [Docker](https://www.docker.com/) image;
  - [Docker Compose](https://docs.docker.com/compose/) configuration.

## [v1.0-alpha](https://github.com/thewizardplusplus/go-exercises-backend/tree/v1.0-alpha) (2021-03-24)

Alpha of the major version. Implement the task model, implement a server; prepare distribution via [Docker](https://www.docker.com/).

### Features

- RESTful API:
  - models:
    - task model:
      - storing:
        - title;
        - description;
        - boilerplate code;
        - test cases:
          - all test cases are represented by a single string;
      - operations:
        - getting all tasks;
        - getting a single task by an ID;
        - creating;
        - updating by an ID;
        - deleting by an ID;
  - representing:
    - in a JSON:
      - payloads:
        - of requests;
        - of responses;
    - as a plain text:
      - errors;
- server:
  - storing settings in environment variables;
  - supporting graceful shutdown;
  - logging:
    - logging requests;
    - logging errors;
  - panics:
    - recovering on panics;
    - logging of panics;
- databases:
  - storing data in the [PostgreSQL](https://www.postgresql.org/) database;
- distributing:
  - [Docker](https://www.docker.com/) image;
  - [Docker Compose](https://docs.docker.com/compose/) configuration.
