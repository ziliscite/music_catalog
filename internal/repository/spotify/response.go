package spotify

var searchResponse = `{
  "tracks": {
    "href": "https://api.spotify.com/v1/search?offset=0&limit=2&query=A%20Little%20Death&type=track&market=US&locale=en-US,en;q%3D0.9,id;q%3D0.8",
    "limit": 2,
    "next": "https://api.spotify.com/v1/search?offset=2&limit=2&query=A%20Little%20Death&type=track&market=US&locale=en-US,en;q%3D0.9,id;q%3D0.8",
    "offset": 0,
    "previous": null,
    "total": 900,
    "items": [
      {
        "album": {
          "album_type": "single",
          "artists": [
            {
              "external_urls": {
                "spotify": "https://open.spotify.com/artist/77SW9BnxLY8rJ0RciFqkHh"
              },
              "href": "https://api.spotify.com/v1/artists/77SW9BnxLY8rJ0RciFqkHh",
              "id": "77SW9BnxLY8rJ0RciFqkHh",
              "name": "The Neighbourhood",
              "type": "artist",
              "uri": "spotify:artist:77SW9BnxLY8rJ0RciFqkHh"
            }
          ],
          "external_urls": {
            "spotify": "https://open.spotify.com/album/5MOCeDoizSpQ4FnpX8VFky"
          },
          "href": "https://api.spotify.com/v1/albums/5MOCeDoizSpQ4FnpX8VFky",
          "id": "5MOCeDoizSpQ4FnpX8VFky",
          "images": [
            {
              "height": 640,
              "width": 640,
              "url": "https://i.scdn.co/image/ab67616d0000b2736492453ee238cd8546c6850e"
            },
            {
              "height": 300,
              "width": 300,
              "url": "https://i.scdn.co/image/ab67616d00001e026492453ee238cd8546c6850e"
            },
            {
              "height": 64,
              "width": 64,
              "url": "https://i.scdn.co/image/ab67616d000048516492453ee238cd8546c6850e"
            }
          ],
          "is_playable": true,
          "name": "Thank You,",
          "release_date": "2012-12-18",
          "release_date_precision": "day",
          "total_tracks": 2,
          "type": "album",
          "uri": "spotify:album:5MOCeDoizSpQ4FnpX8VFky"
        },
        "artists": [
          {
            "external_urls": {
              "spotify": "https://open.spotify.com/artist/77SW9BnxLY8rJ0RciFqkHh"
            },
            "href": "https://api.spotify.com/v1/artists/77SW9BnxLY8rJ0RciFqkHh",
            "id": "77SW9BnxLY8rJ0RciFqkHh",
            "name": "The Neighbourhood",
            "type": "artist",
            "uri": "spotify:artist:77SW9BnxLY8rJ0RciFqkHh"
          }
        ],
        "disc_number": 1,
        "duration_ms": 209706,
        "explicit": false,
        "external_ids": {
          "isrc": "USSM11206332"
        },
        "external_urls": {
          "spotify": "https://open.spotify.com/track/0Ot6e3wYVQQ1Us9PM977jE"
        },
        "href": "https://api.spotify.com/v1/tracks/0Ot6e3wYVQQ1Us9PM977jE",
        "id": "0Ot6e3wYVQQ1Us9PM977jE",
        "is_local": false,
        "is_playable": true,
        "name": "A Little Death",
        "popularity": 69,
        "preview_url": null,
        "track_number": 2,
        "type": "track",
        "uri": "spotify:track:0Ot6e3wYVQQ1Us9PM977jE"
      },
      {
        "album": {
          "album_type": "album",
          "artists": [
            {
              "external_urls": {
                "spotify": "https://open.spotify.com/artist/1DjZI46mVZZZYmmmygRnTw"
              },
              "href": "https://api.spotify.com/v1/artists/1DjZI46mVZZZYmmmygRnTw",
              "id": "1DjZI46mVZZZYmmmygRnTw",
              "name": "Reality Club",
              "type": "artist",
              "uri": "spotify:artist:1DjZI46mVZZZYmmmygRnTw"
            }
          ],
          "external_urls": {
            "spotify": "https://open.spotify.com/album/7gOhCvJD152GWf16fhs7Gp"
          },
          "href": "https://api.spotify.com/v1/albums/7gOhCvJD152GWf16fhs7Gp",
          "id": "7gOhCvJD152GWf16fhs7Gp",
          "images": [
            {
              "height": 640,
              "width": 640,
              "url": "https://i.scdn.co/image/ab67616d0000b2736019b1ddb28634421cc291a0"
            },
            {
              "height": 300,
              "width": 300,
              "url": "https://i.scdn.co/image/ab67616d00001e026019b1ddb28634421cc291a0"
            },
            {
              "height": 64,
              "width": 64,
              "url": "https://i.scdn.co/image/ab67616d000048516019b1ddb28634421cc291a0"
            }
          ],
          "is_playable": true,
          "name": "What Do You Really Know?",
          "release_date": "2019-08-30",
          "release_date_precision": "day",
          "total_tracks": 11,
          "type": "album",
          "uri": "spotify:album:7gOhCvJD152GWf16fhs7Gp"
        },
        "artists": [
          {
            "external_urls": {
              "spotify": "https://open.spotify.com/artist/1DjZI46mVZZZYmmmygRnTw"
            },
            "href": "https://api.spotify.com/v1/artists/1DjZI46mVZZZYmmmygRnTw",
            "id": "1DjZI46mVZZZYmmmygRnTw",
            "name": "Reality Club",
            "type": "artist",
            "uri": "spotify:artist:1DjZI46mVZZZYmmmygRnTw"
          }
        ],
        "disc_number": 1,
        "duration_ms": 238542,
        "explicit": false,
        "external_ids": {
          "isrc": "FR96X1903505"
        },
        "external_urls": {
          "spotify": "https://open.spotify.com/track/6ZGgaShxOimGDfRz1T1zje"
        },
        "href": "https://api.spotify.com/v1/tracks/6ZGgaShxOimGDfRz1T1zje",
        "id": "6ZGgaShxOimGDfRz1T1zje",
        "is_local": false,
        "is_playable": true,
        "name": "Alexandra",
        "popularity": 64,
        "preview_url": null,
        "track_number": 9,
        "type": "track",
        "uri": "spotify:track:6ZGgaShxOimGDfRz1T1zje"
      }
    ]
  }
}`
