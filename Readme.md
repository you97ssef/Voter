# Voter

## Overview

Voter is a small, simple, and secure voting system designed to facilitate the creation, management, and participation in both public and private polls. With a user-friendly interface and robust backend infrastructure, Voter ensures that polls are conducted smoothly and results are trustworthy.

### Live Application

You can access the live application at this URL: [https://voter.bahi.ovh](https://voter.bahi.ovh) and the API at this URL: [https://api.voter.bahi.ovh](https://api.voter.bahi.ovh).

### Key Features

1. **Poll Creation**:

    - **Public Polls**: Easily sharable, open to everyone, allowing broad participation and engagement.
    - **Private Polls**: Secured with unique codes, easily shareable among people who need to participate.
    - **Options**: Each poll must have at least two options to choose from.

2. **Real-time Interaction**:

    - Users can view live updates of poll results as votes are cast.
    - Polls can be closed or deleted by their creators at any time.

3. **Secure Voting**:
    - Users are allowed to vote on only one option per poll to ensure fairness.
    - Guests without accounts can also vote on a poll using a unique name chosen by them.
    - Votes are stored using blockchain technology, providing immutability and security, making it impossible to tamper with the votes without being detected.
    - Users can validate votes of a poll by verifying the hashes of the votes.

## Technology Stack

Voter leverages modern technologies:

-   **Backend**: Developed in Golang, ensuring high performance and concurrency handling.
-   **Database**: Utilizes libsql for efficient and reliable data management.
-   **Frontend**: Built with Vue.js, offering a dynamic and responsive user interface.
-   **Live Updates**: Implemented using MQTT and WebSockets to provide real-time viewing of poll results.
-   **Styling**: Uses Tailwind CSS with Daisy UI configured for a clean and modern look.

## Installation

### Database Setup

For the database, you can use Libsql from your local machine or [Turso](https://turso.tech/), a cloud-based database service.

### Backend Setup

-   Configure the backend by setting up the **environment.json** file in the **api** directory with the right values. You can use the **environment-sample.json** file as a template.

    -   **config**: Configuration about the server and database.
    -   **mailer**: Mail configuration for sending emails with SMTP and credentials.
    -   **jwt**: JWT configuration for generating tokens.
    -   **hasher**: Hash configuration for hashing passwords with cost.

-   You can run the backend using the Docker Compose file in the root of the project. Use the following command to run the backend:

```bash
docker-compose up
```

-   Alternatively, you need to have Golang installed [(Download Golang)](https://go.dev/doc/install) and run the following commands in the **api** directory:

```bash
cd api
go run main.go
```

### Frontend Setup

You need to have Node installed [(Download Node)](https://nodejs.org/en/download/package-manager/current) and run the frontend using the following commands in the **client** directory:

```bash
cd web
npm install
npm run dev
```

### MQTT Broker Setup

-   An online MQTT broker is already set up on the frontend, so you're good to go. If you want to use your own MQTT broker, you can set it up in the frontend by changing the config in the **client/src/main.ts** file.

### Access the Application

-   Open your browser and navigate to the frontend URL to access the Voter application.

## LICENSE

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
