# Groupie-tracker

## Description

Groupie Trackers is a project that involves working with a given API to manipulate the contained data and create a website that displays the information in a user-friendly manner.

- It was given an [API](https://groupietrackers.herokuapp.com/api), that consists in four parts::
    - The first one, `artists`, containing information about some bands and artists like their name(s), image, in which year they began their activity, the date of their first album and the members.
    - The second one, `locations`, consists in their last and/or upcoming concert locations.
    - The third one, `dates`, consists in their last and/or upcoming concert dates.
    - And the last one, `relation`, does the link between all the other parts, `artists`, `dates` and `locations`.

Given all this a user friendly website was built where you can see the bands info through several data visualizations.

The **search** consists of creating a functional program that searches, inside your website, for a specific text input. So the focus of this project is to create a way for the client to search a member or artist or any other attribute in the data system you made.
- The program should handle at least these search cases :
    - artist/band name
    - members
    - locations
    - first album date
    - creation date

The **filter** consists core focus is to empower users with the capability to fine-tune their search outcomes within your website. By implementing effective filter functionality, users can effortlessly sift through data, focus on specific attributes, and access comprehensive details pertaining to members, artists, or other relevant components.


The **geolocalization**  consists on mapping the different concerts locations of a certain artist/band given by the Client.

## Usage

1. Run server

```
$ go run cmd/main.go
```

Server will start at **http://localhost:8081** address


