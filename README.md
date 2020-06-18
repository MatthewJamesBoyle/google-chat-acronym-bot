
# google-chat-acronym-bot
(A sample version of this project is currently running over at https://gentle-springs-50534.herokuapp.com)

This is a really simple implementation of a chatbot to help users dispel the many acronyms your company has accured over the years. It supports the following commands:

- add {acronym} {definition} 
- explain {acronym} 
- help

Its currently wired up to use an "in memory" database (really its just a hash map) but its interfaced well so adding a "real" db would be fairly simple.

Although I wrote this for google chat, it could be very easily adopted for any of the other platforms (slack, discord etc) by simply making some minor tweaks to  `main.go`

## Setting it up For Google Chat
See the instructions [here](https://developers.google.com/hangouts/chat). You basically just need to deploy it and enter the url in your google dashboard.
**Note: The implementation currently does not have any auth. I highly recommend you add it before using it in any meaningful way.**

## Deploying it to Heroku.
This repo contains everything you need to deploy it to your own Heroku account.
Make sure you have the Heroku CLI installed, and then run:

 - `Heroku login`
 - `make Heroku`
 - `heroku container release:web`
Thats it!
