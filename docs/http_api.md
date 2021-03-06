# HTTP API

This document contains design for HTTP API used in this system.

Everytime client want to call this API, it need to embed `X-API-Key: c4211664-47dc-4887-a2fe-9e694fbaf55a` in request header.

## New Game

POST: `/games`

This endpoint is used for starting new game. The initial scenario for the game is `NEW_QUESTION` which tell client that it should fetch new question to advance the scenario.

**Body Fields:**

- `player_name`, String => name of player that play the game.

**Example Request:**

```json
POST /games
Content-Type: application/json

{
    "player_name": "Riandy"
}
```

**Success Response:**

```json
HTTP/1.1 200 OK
Content-Type: application/json

{
    "ok": true,
    "data": {
        "game_id": "7bb9cd49-bff1-45f3-982d-e533c8336989",
        "player_name": "Riandy",
        "scenario": "NEW_QUESTION"
    },
    "ts": 1648814458
}
```

**Error Responses:**

- Bad Request (`404`)

    ```json
    HTTP/1.1 404 Bad Request
    Content-Type: application/json

    {
        "ok": false,
        "err": "ERR_BAD_REQUEST",
        "msg": "missing `player_name`",
        "ts": 1648814458
    }
    ```


[Back to Top](#http-api)

---

## New Question

PUT: `/games/{game_id}/question`

This endpoint is used for generating new question. After successfully calling it, the game scenario is set to `SUBMIT_ANSWER` which telling client it should submit answer in the next call.

The success response includes `timeout_at` which tell the time when the question is timed out. When client submit the answer during or after `timeout_at` the game will end immediately.

**Example Request:**

```json
PUT /games/7bb9cd49-bff1-45f3-982d-e533c8336989/question
```

**Success Response:**

```json
HTTP/1.1 200 OK
Content-Type: application/json

{
    "ok": true,
    "data": {
        "game_id": "7bb9cd49-bff1-45f3-982d-e533c8336989",
        "scenario": "SUBMIT_ANSWER",
        "problem": "1 + 2",
        "choices": [
            "3",
            "4",
            "5"
        ],
        "timeout": 5
    },
    "ts": 1648814458
}
```

**Error Responses:**

- Invalid Scenario (`409`)

    ```json
    HTTP/1.1 409 Conflict
    Content-Type: application/json

    {
        "ok": false,
        "err": "ERR_INVALID_SCENARIO",
        "msg": "invalid scenario for the action",
        "ts": 1648814458
    }
    ```

    Client will receive this error when scenario is not `NEW_QUESTION` upon calling this endpoint.

[Back to Top](#http-api)

---

## Submit Answer

PUT: `/games/{game_id}/answer`

This endpoint is used for sending answer from client.

Notice that even though user don't submit any answer, this endpoint still need to be called, otherwise the game will stale. From the backend perspective this is pretty much okay, but if in the future if we decided to add leaderboard to this game, the staled game will resulted in user score doesn't submitted to leaderboard. So it's better to always call this endpoint whether user submitting answer or not.

There are 2 next possible scenarios after successfully calling this endpoint:

- `NEW_QUESTION` => if the submitted answer is correct
- `GAME_OVER` => if either question has been timed out or submitted answer is incorrect 

**Body Fields:**

- `answer_idx`, Integer => answer index for the question, start from `1`, if user doesn't give any answer set the value to `0`.
- `start_at`, Integer => unix timestamp in seconds when client start displaying the question.
- `sent_at`, Integer => unix timestamp in seconds when the answer was sent, the value of this timestamp will be used to determine whether or not user submit the answer during or after timeout.

**Example Request:**

```json
PUT /games/7bb9cd49-bff1-45f3-982d-e533c8336989/answer
Content-Type: application/json

{
    "answer_idx": 1,
    "start_at": 1648814453,
    "sent_at": 1648814458
}
```

**Success Responses:**

- Correct answer:

    ```json
    HTTP/1.1 200 OK
    Content-Type: application/json

    {
        "ok": true,
        "data": {
            "game_id": "7bb9cd49-bff1-45f3-982d-e533c8336989",
            "scenario": "NEW_QUESTION",
            "answer_idx": 1,
            "correct_idx": 1,
            "duration": 4, // the duration of user answering question, calculated from sent_at - start_at
            "timeout": 5, // the timeout duration given by server
            "score": 20,
        },
        "ts": 1648814458
    }
    ```

- Incorrect answer or time out:

    ```json
    HTTP/1.1 200 OK
    Content-Type: application/json

    {
        "ok": true,
        "data": {
            "game_id": "7bb9cd49-bff1-45f3-982d-e533c8336989",
            "scenario": "GAME_OVER",
            "answer_idx": 1,
            "correct_idx": 2,
            "duration": 4,
            "timeout": 5,
            "score": 0
        },
        "ts": 1648814458
    }
    ```

**Error Responses:**

- Bad Request (`404`)

    ```json
    HTTP/1.1 400 Bad Request
    Content-Type: application/json

    {
        "ok": false,
        "err": "ERR_BAD_REQUEST",
        "msg": "invalid value of `sent_at`"
    }
    ```

    Client will receive error like this when the parameter value in the request is invalid.

- Invalid Scenario (`409`)

    ```json
    HTTP/1.1 409 Conflict
    Content-Type: application/json

    {
        "ok": false,
        "err": "ERR_INVALID_SCENARIO",
        "msg": "invalid scenario for the action",
        "ts": 1648814458
    }
    ```

    Client will receive this error when scenario is not `SUBMIT_ANSWER` upon calling this endpoint.

[Back to Top](#http-api)

---