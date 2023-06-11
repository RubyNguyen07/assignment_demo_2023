# assignment_demo_2023

![Tests](https://github.com/TikTokTechImmersion/assignment_demo_2023/actions/workflows/test.yml/badge.svg)

This is my implementation for backend assignment of 2023 TikTok Tech Immersion.

Requirements: https://bytedance.sg.feishu.cn/docx/P9kQdDkh5oqG37xVm5slN1Mrgle

## How to run with Docker 
Make sure you have Docker installed:
- Windows: https://docs.docker.com/desktop/install/windows-install/
- MacOS: https://docs.docker.com/desktop/install/mac-install/
Right-click docker-compose.yml and run

## API Documentation 
<ol>
    <li> Ping: Check if the server is running </li>
    ```
        curl -X GET http://localhost:8080/ping
    ```
    Expected response: status 200 
    ```json
        {
            "message": "pong"
        }
    ```
    <li> Send message: send message in a chat room </li>
    ```
        curl -X POST \
          http://localhost:8080/api/send \
          -H 'Content-Type: application/json' \
          -d '{
            "Chat": "jenny:lisa",
            "Text": "Hello World",
            "Sender": "jenny"
        }'
    ```
    Expected response: status 200
    <li> Pull messages: retrieve messages in a chat room from Cursor with Limit and sorting order Reverse (default: False)</li>
    ```
        curl -X GET \
          http://localhost:8080/api/pull \
          -H 'Content-Type: application/json' \
          -d '{
            "Chat": "jenny:lisa",
            "Cursor": 0,
            "Limit": 20,
            "Reverse": false
        }'
    ```
    Expected response: status 200
    {
      "messages": [
        {
          "chat": "jenny:lisa",
          "text": "Hello World",
          "sender": "jenny",
          "send_time": 1684744610
        }, ...
      ]
    }
    > Note that the order of members in a chat is not important 
</ol> 

## Tech stack 
- Golang 
- Redis
- Kitex
- Docker 
