# Change Log

## [v1.7.2](https://github.com/thewizardplusplus/go-exercises-backend/tree/v1.7.2) (2022-08-28)

Add a few example tasks.

- example tasks:
  - add the migration that creates a few example tasks:
    - add the utility for applying migrations;
  - add the solutions for the example tasks.

## [v1.7.1](https://github.com/thewizardplusplus/go-exercises-backend/tree/v1.7.1) (2022-05-23)

Perform refactoring.

- perform refactoring:
  - replace the `registers.SolutionResultRegister` structure with the `usecases.SolutionUsecase.RegisterSolutionResult()` method;
  - replace the `registers.SolutionRegister.performRegistration()` method with the `usecases.SolutionUsecase.RegisterSolution()` one;
  - add the `usecases.UserUsecase` structure.

## [v1.7](https://github.com/thewizardplusplus/go-exercises-backend/tree/v1.7) (2022-05-21)

Support for user updating in the utility for managing users.

- utilities:
  - utility for managing users:
    - commands:
      - add a user:
        - parameters:
          - flag indicating the need to generate a password;
          - flag indicating whether the user is disabled or not;
        - simplify control of password generation;
      - update a user:
        - parameters:
          - username;
          - new username;
          - password;
          - password hashing cost;
          - flag indicating the need to generate a password;
          - generated password length;
          - flag indicating whether the user is disabled or not;
          - flag indicating whether the user is enabled or not;
        - can update the user fields individually.

### Features

- RESTful API:
  - models:
    - task model:
      - storing:
        - author ID;
        - title;
        - description;
        - boilerplate code;
        - test cases:
          - all test cases are represented by a single string;
      - operations:
        - getting all tasks:
          - calculate a total correctness flag based on all solutions of a requesting author;
          - sort the results by creation time in descending order;
          - process pagination:
            - implemented using an offset and a limit;
        - getting a single task by an ID:
          - calculate a total correctness flag based on all solutions of a requesting author;
        - creating:
          - automatically format a task boilerplate code;
        - updating by an ID:
          - allowed for its author only;
          - automatically format a task boilerplate code;
        - deleting by an ID:
          - allowed for its author only;
    - solution model:
      - storing:
        - author ID;
        - task ID;
        - code;
        - correctness flag;
        - testing result:
          - represented by a string;
      - operations:
        - getting all solutions by a task ID:
          - filtered by a requesting author:
            - allow a solution task author to get solutions of other authors;
          - sort the results by creation time in descending order;
          - process pagination:
            - implemented using an offset and a limit;
        - getting a single solution by an ID:
          - allowed for:
            - solution author;
            - solution task author;
        - creating:
          - automatically format a solution code;
        - updating by an ID:
          - performed by a queue consumer only (see below);
        - formatting a solution code;
  - representing:
    - in a JSON:
      - payloads:
        - of requests;
        - of responses;
    - as a plain text:
      - errors;
- server:
  - additional routing:
    - serving static files;
  - storing settings in environment variables;
  - supporting graceful shutdown;
  - logging:
    - logging requests;
    - logging errors;
  - panics:
    - recovering on panics;
    - logging of panics;
- authentication:
  - use the Bearer authentication scheme based on [JSON Web Tokens](https://jwt.io/):
    - store in a JWT claims:
      - expiration time claim;
      - user claim:
        - contains a whole user model;
  - generate a token signing key automatically by default;
  - user model:
    - storing:
      - username (unique);
      - password hash:
        - generated using the [bcrypt](https://en.wikipedia.org/wiki/Bcrypt) function;
      - flag indicating whether the user is disabled or not;
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
    - consuming solution results:
      - concurrent consuming;
      - once requeue the solution on failure;
- utilities:
  - utility for managing users:
    - commands:
      - add a user:
        - parameters:
          - username;
          - password;
          - password hashing cost;
          - flag indicating the need to generate a password;
          - generated password length;
          - flag indicating whether the user is disabled or not;
      - update a user:
        - parameters:
          - username;
          - new username;
          - password;
          - password hashing cost;
          - flag indicating the need to generate a password;
          - generated password length;
          - flag indicating whether the user is disabled or not;
          - flag indicating whether the user is enabled or not;
        - can update the user fields individually;
- distributing:
  - [Docker](https://www.docker.com/) image;
  - [Docker Compose](https://docs.docker.com/compose/) configuration.

## [v1.6](https://github.com/thewizardplusplus/go-exercises-backend/tree/v1.6) (2022-03-24)

Support for disabling access to specific users.

- authentication:
  - user model:
    - storing:
      - flag indicating whether the user is disabled or not.

### Features

- RESTful API:
  - models:
    - task model:
      - storing:
        - author ID;
        - title;
        - description;
        - boilerplate code;
        - test cases:
          - all test cases are represented by a single string;
      - operations:
        - getting all tasks:
          - calculate a total correctness flag based on all solutions of a requesting author;
          - sort the results by creation time in descending order;
          - process pagination:
            - implemented using an offset and a limit;
        - getting a single task by an ID:
          - calculate a total correctness flag based on all solutions of a requesting author;
        - creating:
          - automatically format a task boilerplate code;
        - updating by an ID:
          - allowed for its author only;
          - automatically format a task boilerplate code;
        - deleting by an ID:
          - allowed for its author only;
    - solution model:
      - storing:
        - author ID;
        - task ID;
        - code;
        - correctness flag;
        - testing result:
          - represented by a string;
      - operations:
        - getting all solutions by a task ID:
          - filtered by a requesting author:
            - allow a solution task author to get solutions of other authors;
          - sort the results by creation time in descending order;
          - process pagination:
            - implemented using an offset and a limit;
        - getting a single solution by an ID:
          - allowed for:
            - solution author;
            - solution task author;
        - creating:
          - automatically format a solution code;
        - updating by an ID:
          - performed by a queue consumer only (see below);
        - formatting a solution code;
  - representing:
    - in a JSON:
      - payloads:
        - of requests;
        - of responses;
    - as a plain text:
      - errors;
- server:
  - additional routing:
    - serving static files;
  - storing settings in environment variables;
  - supporting graceful shutdown;
  - logging:
    - logging requests;
    - logging errors;
  - panics:
    - recovering on panics;
    - logging of panics;
- authentication:
  - use the Bearer authentication scheme based on [JSON Web Tokens](https://jwt.io/):
    - store in a JWT claims:
      - expiration time claim;
      - user claim:
        - contains a whole user model;
  - generate a token signing key automatically by default;
  - user model:
    - storing:
      - username (unique);
      - password hash:
        - generated using the [bcrypt](https://en.wikipedia.org/wiki/Bcrypt) function;
      - flag indicating whether the user is disabled or not;
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
    - consuming solution results:
      - concurrent consuming;
      - once requeue the solution on failure;
- utilities:
  - utility for managing users:
    - commands:
      - add a user:
        - parameters:
          - username;
          - password:
            - generate automatically by default;
          - password hashing cost;
          - generated password length;
- distributing:
  - [Docker](https://www.docker.com/) image;
  - [Docker Compose](https://docs.docker.com/compose/) configuration.

## [v1.5.3](https://github.com/thewizardplusplus/go-exercises-backend/tree/v1.5.3) (2022-03-24)

Describe the releases of the project.

- describe for the releases of the project:
  - features;
  - change log;
- utilities:
  - utility for managing users:
    - add the `README.md` file.

## [v1.5.2](https://github.com/thewizardplusplus/go-exercises-backend/tree/v1.5.2) (2022-02-27)

Describe the RabbitMQ API and the web API.

- describe the RabbitMQ API in the [AsyncAPI](https://www.asyncapi.com/) format;
- describe the web API in the [Swagger](http://swagger.io/) format.

## [v1.5.1](https://github.com/thewizardplusplus/go-exercises-backend/tree/v1.5.1) (2021-11-17)

Perform refactoring; improve the HTTP statuses of the responses.

- improve the HTTP statuses of the responses:
  - in case of success;
  - in case of an error;
- perform refactoring:
  - add the `usecases` package.

## [v1.5](https://github.com/thewizardplusplus/go-exercises-backend/tree/v1.5) (2021-11-17)

Perform refactoring; additionally return total result count along with pagination results.

- RESTful API:
  - models:
    - task model:
      - operations:
        - getting all tasks:
          - process pagination:
            - additionally return total result count;
    - solution model:
      - operations:
        - getting all solutions by a task ID:
          - process pagination:
            - additionally return total result count;
        - getting a single solution by an ID:
          - join a solution task data to the results;
- perform refactoring:
  - update the [github.com/thewizardplusplus/go-http-utils](https://github.com/thewizardplusplus/go-http-utils) package;
  - add the [github.com/thewizardplusplus/go-rabbitmq-utils](https://github.com/thewizardplusplus/go-rabbitmq-utils) package;
  - add the [github.com/thewizardplusplus/go-sync-utils](https://github.com/thewizardplusplus/go-sync-utils) package;
- utilities:
  - utility for managing users:
    - add to the [Docker](https://www.docker.com/) image.

## [v1.4](https://github.com/thewizardplusplus/go-exercises-backend/tree/v1.4) (2021-05-10)

Format a task boilerplate code and a solution code (including automatically).

- RESTful API:
  - models:
    - task model:
      - operations:
        - creating:
          - automatically format a task boilerplate code;
        - updating by an ID:
          - automatically format a task boilerplate code;
    - solution model:
      - operations:
        - creating:
          - automatically format a solution code;
        - formatting a solution code.

### Features

- RESTful API:
  - models:
    - task model:
      - storing:
        - author ID;
        - title;
        - description;
        - boilerplate code;
        - test cases:
          - all test cases are represented by a single string;
      - operations:
        - getting all tasks:
          - calculate a total correctness flag based on all solutions of a requesting author;
          - sort the results by creation time in descending order;
          - process pagination:
            - implemented using an offset and a limit;
        - getting a single task by an ID:
          - calculate a total correctness flag based on all solutions of a requesting author;
        - creating:
          - automatically format a task boilerplate code;
        - updating by an ID:
          - allowed for its author only;
          - automatically format a task boilerplate code;
        - deleting by an ID:
          - allowed for its author only;
    - solution model:
      - storing:
        - author ID;
        - task ID;
        - code;
        - correctness flag;
        - testing result:
          - represented by a string;
      - operations:
        - getting all solutions by a task ID:
          - filtered by a requesting author:
            - allow a solution task author to get solutions of other authors;
          - sort the results by creation time in descending order;
          - process pagination:
            - implemented using an offset and a limit;
        - getting a single solution by an ID:
          - allowed for:
            - solution author;
            - solution task author;
        - creating:
          - automatically format a solution code;
        - updating by an ID:
          - performed by a queue consumer only (see below);
        - formatting a solution code;
  - representing:
    - in a JSON:
      - payloads:
        - of requests;
        - of responses;
    - as a plain text:
      - errors;
- server:
  - additional routing:
    - serving static files;
  - storing settings in environment variables;
  - supporting graceful shutdown;
  - logging:
    - logging requests;
    - logging errors;
  - panics:
    - recovering on panics;
    - logging of panics;
- authentication:
  - use the Bearer authentication scheme based on [JSON Web Tokens](https://jwt.io/):
    - store in a JWT claims:
      - expiration time claim;
      - user claim:
        - contains a whole user model;
  - generate a token signing key automatically by default;
  - user model:
    - storing:
      - username (unique);
      - password hash:
        - generated using the [bcrypt](https://en.wikipedia.org/wiki/Bcrypt) function;
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
    - consuming solution results:
      - concurrent consuming;
      - once requeue the solution on failure;
- utilities:
  - utility for managing users:
    - commands:
      - add a user:
        - parameters:
          - username;
          - password:
            - generate automatically by default;
          - password hashing cost;
          - generated password length;
- distributing:
  - [Docker](https://www.docker.com/) image;
  - [Docker Compose](https://docs.docker.com/compose/) configuration.

## [v1.3](https://github.com/thewizardplusplus/go-exercises-backend/tree/v1.3) (2021-05-08)

Calculate a total correctness flag based on all solutions of a requesting author; sort the tasks and solutions; process pagination for the tasks and solutions; allow a solution task author to get solutions of other authors.

- RESTful API:
  - models:
    - task model:
      - operations:
        - getting all tasks:
          - join an additional data:
            - join an author data to the results;
            - calculate a total correctness flag based on all solutions of a requesting author;
          - sort the results by creation time in descending order;
          - process pagination:
            - implemented using an offset and a limit;
        - getting a single task by an ID:
          - join an additional data:
            - join an author data to the results;
            - calculate a total correctness flag based on all solutions of a requesting author;
    - solution model:
      - operations:
        - getting all solutions by a task ID:
          - join an author data to the results;
          - filtered by a requesting author:
            - allow a solution task author to get solutions of other authors;
          - sort the results by creation time in descending order;
          - process pagination:
            - implemented using an offset and a limit;
        - getting a single solution by an ID:
          - join an author data to the results;
          - allow a solution task author to get solutions of other authors.

### Features

- RESTful API:
  - models:
    - task model:
      - storing:
        - author ID;
        - title;
        - description;
        - boilerplate code;
        - test cases:
          - all test cases are represented by a single string;
      - operations:
        - getting all tasks:
          - calculate a total correctness flag based on all solutions of a requesting author;
          - sort the results by creation time in descending order;
          - process pagination:
            - implemented using an offset and a limit;
        - getting a single task by an ID:
          - calculate a total correctness flag based on all solutions of a requesting author;
        - creating;
        - updating by an ID:
          - allowed for its author only;
        - deleting by an ID:
          - allowed for its author only;
    - solution model:
      - storing:
        - author ID;
        - task ID;
        - code;
        - correctness flag;
        - testing result:
          - represented by a string;
      - operations:
        - getting all solutions by a task ID:
          - filtered by a requesting author:
            - allow a solution task author to get solutions of other authors;
          - sort the results by creation time in descending order;
          - process pagination:
            - implemented using an offset and a limit;
        - getting a single solution by an ID:
          - allowed for:
            - solution author;
            - solution task author;
        - creating;
        - updating by an ID:
          - performed by a queue consumer only (see below);
  - representing:
    - in a JSON:
      - payloads:
        - of requests;
        - of responses;
    - as a plain text:
      - errors;
- server:
  - additional routing:
    - serving static files;
  - storing settings in environment variables;
  - supporting graceful shutdown;
  - logging:
    - logging requests;
    - logging errors;
  - panics:
    - recovering on panics;
    - logging of panics;
- authentication:
  - use the Bearer authentication scheme based on [JSON Web Tokens](https://jwt.io/):
    - store in a JWT claims:
      - expiration time claim;
      - user claim:
        - contains a whole user model;
  - generate a token signing key automatically by default;
  - user model:
    - storing:
      - username (unique);
      - password hash:
        - generated using the [bcrypt](https://en.wikipedia.org/wiki/Bcrypt) function;
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
    - consuming solution results:
      - concurrent consuming;
      - once requeue the solution on failure;
- utilities:
  - utility for managing users:
    - commands:
      - add a user:
        - parameters:
          - username;
          - password:
            - generate automatically by default;
          - password hashing cost;
          - generated password length;
- distributing:
  - [Docker](https://www.docker.com/) image;
  - [Docker Compose](https://docs.docker.com/compose/) configuration.

## [v1.2](https://github.com/thewizardplusplus/go-exercises-backend/tree/v1.2) (2021-04-02)

Implement serving static files.

- server:
  - additional routing:
    - serving static files.

### Features

- RESTful API:
  - models:
    - task model:
      - storing:
        - author ID;
        - title;
        - description;
        - boilerplate code;
        - test cases:
          - all test cases are represented by a single string;
      - operations:
        - getting all tasks;
        - getting a single task by an ID;
        - creating;
        - updating by an ID:
          - allowed for its author only;
        - deleting by an ID:
          - allowed for its author only;
    - solution model:
      - storing:
        - author ID;
        - task ID;
        - code;
        - correctness flag;
        - testing result:
          - represented by a string;
      - operations:
        - getting all solutions by a task ID:
          - filtered by a requesting author;
        - getting a single solution by an ID:
          - allowed for its author only;
        - creating;
        - updating by an ID:
          - performed by a queue consumer only (see below);
  - representing:
    - in a JSON:
      - payloads:
        - of requests;
        - of responses;
    - as a plain text:
      - errors;
- server:
  - additional routing:
    - serving static files;
  - storing settings in environment variables;
  - supporting graceful shutdown;
  - logging:
    - logging requests;
    - logging errors;
  - panics:
    - recovering on panics;
    - logging of panics;
- authentication:
  - use the Bearer authentication scheme based on [JSON Web Tokens](https://jwt.io/):
    - store in a JWT claims:
      - expiration time claim;
      - user claim:
        - contains a whole user model;
  - generate a token signing key automatically by default;
  - user model:
    - storing:
      - username (unique);
      - password hash:
        - generated using the [bcrypt](https://en.wikipedia.org/wiki/Bcrypt) function;
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
    - consuming solution results:
      - concurrent consuming;
      - once requeue the solution on failure;
- utilities:
  - utility for managing users:
    - commands:
      - add a user:
        - parameters:
          - username;
          - password:
            - generate automatically by default;
          - password hashing cost;
          - generated password length;
- distributing:
  - [Docker](https://www.docker.com/) image;
  - [Docker Compose](https://docs.docker.com/compose/) configuration.

## [v1.1](https://github.com/thewizardplusplus/go-exercises-backend/tree/v1.1) (2021-04-01)

Implement authentication that use the Bearer authentication scheme based on [JSON Web Tokens](https://jwt.io/); implement the utility for adding users.

- RESTful API:
  - models:
    - task model:
      - storing:
        - author ID;
      - operations:
        - updating by an ID:
          - allowed for its author only;
        - deleting by an ID:
          - allowed for its author only;
    - solution model:
      - storing:
        - author ID;
      - operations:
        - getting all solutions by a task ID:
          - filtered by a requesting author;
        - getting a single solution by an ID:
          - allowed for its author only;
- authentication:
  - use the Bearer authentication scheme based on [JSON Web Tokens](https://jwt.io/):
    - store in a JWT claims:
      - expiration time claim;
      - user claim:
        - contains a whole user model;
  - generate a token signing key automatically by default;
  - user model:
    - storing:
      - username (unique);
      - password hash:
        - generated using the [bcrypt](https://en.wikipedia.org/wiki/Bcrypt) function;
- utilities:
  - utility for managing users:
    - commands:
      - add a user:
        - parameters:
          - username;
          - password:
            - generate automatically by default;
          - password hashing cost;
          - generated password length.

### Features

- RESTful API:
  - models:
    - task model:
      - storing:
        - author ID;
        - title;
        - description;
        - boilerplate code;
        - test cases:
          - all test cases are represented by a single string;
      - operations:
        - getting all tasks;
        - getting a single task by an ID;
        - creating;
        - updating by an ID:
          - allowed for its author only;
        - deleting by an ID:
          - allowed for its author only;
    - solution model:
      - storing:
        - author ID;
        - task ID;
        - code;
        - correctness flag;
        - testing result:
          - represented by a string;
      - operations:
        - getting all solutions by a task ID:
          - filtered by a requesting author;
        - getting a single solution by an ID:
          - allowed for its author only;
        - creating;
        - updating by an ID:
          - performed by a queue consumer only (see below);
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
- authentication:
  - use the Bearer authentication scheme based on [JSON Web Tokens](https://jwt.io/):
    - store in a JWT claims:
      - expiration time claim;
      - user claim:
        - contains a whole user model;
  - generate a token signing key automatically by default;
  - user model:
    - storing:
      - username (unique);
      - password hash:
        - generated using the [bcrypt](https://en.wikipedia.org/wiki/Bcrypt) function;
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
    - consuming solution results:
      - concurrent consuming;
      - once requeue the solution on failure;
- utilities:
  - utility for managing users:
    - commands:
      - add a user:
        - parameters:
          - username;
          - password:
            - generate automatically by default;
          - password hashing cost;
          - generated password length;
- distributing:
  - [Docker](https://www.docker.com/) image;
  - [Docker Compose](https://docs.docker.com/compose/) configuration.

## [v1.0.2](https://github.com/thewizardplusplus/go-exercises-backend/tree/v1.0.2) (2021-03-29)

Extend the [Postman](https://www.postman.com/) collection.

- extend the [Postman](https://www.postman.com/) collection.

## [v1.0.1](https://github.com/thewizardplusplus/go-exercises-backend/tree/v1.0.1) (2021-03-27)

Use the `datatypes.JSON` type from the `github.com/go-gorm/datatypes` package in the models; requeue the solution on failure only once.

- RESTful API:
  - models:
    - task model:
      - storing:
        - test cases:
          - use the `datatypes.JSON` type from the [github.com/go-gorm/datatypes](https://github.com/go-gorm/datatypes) package;
    - solution model:
      - storing:
        - testing result:
          - use the `datatypes.JSON` type from the [github.com/go-gorm/datatypes](https://github.com/go-gorm/datatypes) package;
- interaction with queues:
  - operations:
    - consuming solution results:
      - once requeue the solution on failure.

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
        - correctness flag;
        - testing result:
          - represented by a string;
      - operations:
        - getting all solutions by a task ID;
        - getting a single solution by an ID;
        - creating;
        - updating by an ID;
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
    - consuming solution results:
      - concurrent consuming;
      - once requeue the solution on failure;
- distributing:
  - [Docker](https://www.docker.com/) image;
  - [Docker Compose](https://docs.docker.com/compose/) configuration.

## [v1.0](https://github.com/thewizardplusplus/go-exercises-backend/tree/v1.0) (2021-03-26)

Major version. Implement consuming solution results via the [RabbitMQ](https://www.rabbitmq.com/) message broker.

- RESTful API:
  - models:
    - solution model:
      - storing:
        - correctness flag;
        - testing result:
          - represented by a string;
      - operations:
        - updating by an ID;
- interaction with queues:
  - operations:
    - consuming solution results:
      - concurrent consuming;
      - requeue the solution on failure;
  - set a maximal size for the used queues.

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
        - correctness flag;
        - testing result:
          - represented by a string;
      - operations:
        - getting all solutions by a task ID;
        - getting a single solution by an ID;
        - creating;
        - updating by an ID;
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
    - consuming solution results:
      - concurrent consuming;
      - requeue the solution on failure;
- distributing:
  - [Docker](https://www.docker.com/) image;
  - [Docker Compose](https://docs.docker.com/compose/) configuration.

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
