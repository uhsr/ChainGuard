# ChainGuard: Decentralized Crypto Price Alerts

Real-time, on-chain event driven crypto price alerts delivered via WebSockets and gRPC microservices.

ChainGuard provides a decentralized and highly scalable solution for monitoring cryptocurrency prices and receiving immediate notifications when specific events occur on-chain. Unlike centralized price alert services, ChainGuard directly monitors blockchain events, eliminating reliance on potentially manipulated or delayed data feeds. This system leverages on-chain event listeners to detect price fluctuations, volume changes, and other critical indicators, forwarding these events through a robust gRPC microservice architecture to deliver real-time alerts to users via WebSocket connections. ChainGuard enables traders, developers, and enthusiasts to react instantly to market changes, automate trading strategies, and gain a competitive edge.

The core of ChainGuard lies in its ability to adapt to different blockchains and notification methods. The system is designed with modularity in mind, allowing for easy integration of new blockchain event listeners and notification channels. This flexibility extends to the alert criteria themselves, which can be customized to meet specific user requirements. Whether monitoring for a specific price threshold, a significant trading volume surge, or a large transaction on a particular smart contract, ChainGuard provides the tools necessary to define and receive timely alerts.

By utilizing a microservice architecture, ChainGuard achieves high availability and scalability. The individual services responsible for event listening, alert processing, and notification delivery can be independently scaled to handle increasing workloads. This ensures that the system remains responsive and reliable, even during periods of high market volatility. The gRPC framework enables efficient communication between these services, minimizing latency and maximizing throughput. Furthermore, the decentralized nature of the system reduces the risk of single points of failure, enhancing overall robustness.

## Key Features

*   **On-Chain Event Listeners:** Monitors blockchain events directly using Go's `ethereum/go-ethereum` library for accurate and real-time price detection. Supports multiple blockchains via configurable event listeners.
*   **gRPC Microservices:** Utilizes gRPC for inter-service communication, ensuring high performance and scalability. Services include: `event_listener`, `alert_processor`, and `notification_service`.
*   **WebSocket Notifications:** Delivers real-time price alerts to users through persistent WebSocket connections, leveraging Go's `gorilla/websocket` package.
*   **Customizable Alert Rules:** Allows users to define specific alert criteria based on price thresholds, volume changes, and other on-chain events. Alert rules are stored in a database and dynamically loaded by the `alert_processor`.
*   **Configurable Notification Channels:** Supports multiple notification channels, including WebSocket, email (via SMTP), and push notifications (via third-party services).
*   **Decentralized Architecture:** Eliminates reliance on centralized price feeds, improving accuracy and transparency.
*   **Scalable Design:** The microservice architecture allows for independent scaling of individual components to handle increasing workloads.

## Technology Stack

*   **Go:** The primary programming language used for all services.
*   **gRPC:** Used for inter-service communication, providing high performance and scalability.
*   **Protocol Buffers:** Used for defining the gRPC service contracts and message formats.
*   **Ethereum/go-ethereum:** Used for interacting with the Ethereum blockchain and listening for events.
*   **Gorilla/websocket:** Used for creating and managing WebSocket connections for real-time notifications.
*   **PostgreSQL:** Used for storing alert rules and user preferences.
*   **Docker:** Used for containerizing the services for easy deployment and management.

## Installation

1.  **Clone the repository:**

    git clone https://github.com/uhsr/ChainGuard.git
    cd ChainGuard

2.  **Install Go dependencies:**

    go mod download
    go mod vendor

3.  **Install Docker and Docker Compose:** Follow the instructions for your operating system.

4.  **Configure environment variables:** Create a `.env` file in the root directory based on the `.env.example` file.

5.  **Build the Docker images:**

    docker-compose build

6.  **Start the services:**

    docker-compose up -d

## Configuration

The application relies on environment variables for configuration. The following environment variables are required:

*   `DATABASE_URL`: PostgreSQL connection string.
*   `ETHEREUM_NODE_URL`: URL of the Ethereum node to connect to.
*   `WS_PORT`: Port for the WebSocket server.
*   `GRPC_PORT`: Port for the gRPC server.
*   `SMTP_HOST`: SMTP server hostname for email notifications.
*   `SMTP_PORT`: SMTP server port for email notifications.
*   `SMTP_USER`: SMTP server username for email notifications.
*   `SMTP_PASSWORD`: SMTP server password for email notifications.

## Usage

After installation, the services will be running in Docker containers. The `event_listener` service will start listening for events on the configured Ethereum node. The `alert_processor` service will periodically query the database for new alert rules and process incoming events. The `notification_service` will deliver alerts to users via WebSocket connections.

To connect to the WebSocket server, use a WebSocket client and connect to `ws://localhost:<WS_PORT>`.

To interact with the gRPC services, use a gRPC client and connect to `localhost:<GRPC_PORT>`. The Protocol Buffer definitions for the gRPC services are located in the `proto` directory.

## Contributing

We welcome contributions to ChainGuard! Please follow these guidelines:

*   Fork the repository and create a branch for your changes.
*   Write clear and concise commit messages.
*   Include unit tests for your changes.
*   Submit a pull request to the `main` branch.

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/uhsr/ChainGuard/blob/main/LICENSE) file for details.

## Acknowledgements

We would like to thank the open-source community for their contributions to the Go ecosystem.