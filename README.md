# ActivityStream 🏞

[![GoDoc](https://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://pkg.go.dev/github.com/benpate/activitystream)
[![Build Status](https://img.shields.io/github/workflow/status/benpate/pub/Go/main)](https://github.com/benpate/activitystream/actions/workflows/go.yml)
[![Codecov](https://img.shields.io/codecov/c/github/benpate/pub.svg?style=flat-square)](https://codecov.io/gh/benpate/pub)
[![Go Report Card](https://goreportcard.com/badge/github.com/benpate/activitystream?style=flat-square)](https://goreportcard.com/report/github.com/benpate/activitystream)
[![Version](https://img.shields.io/github/v/release/benpate/pub?include_prereleases&style=flat-square&color=brightgreen)](https://github.com/benpate/activitystream/releases)

## Welcome to the Fediverse

This is a personal experiment to explore Activity Streams.  This library is BRAND NEW, and is not ready for use **by anyone**, for **any reason**, at **any time**.  Check out [https://github.com/go-fed](https://github.com/go-fed) for a library that's ready for prime time.

## This is an Experiment

ActivityStreams is so hard to work with in a "strongly typed", "idiomatic go" way, because the W3C spec is so loose with types.  While this is *great* for a Javascript, or an old ColdFusion programmer, it's super-cumbersome to try to squeeze this into a Go, or TypeScript paradigm.

So, here's the idea (for now).  EVERYTHING is *stored* as a `map[string]interface{}`, which is as big of a no-no as using Reflect.  But, it's the only way to support the entire spec without doing [this](https://github.com/go-fed/activity/blob/master/streams/vocab/gen_type_activitystreams_accept_interface.go).

Consistency will be "enforced" (or at least *encouraged*) using constructor functions that can construct these data structures in a "proper" way, and validator functions that can accept (nearly) any input, and then format it in a consistent manner.

Something like:

```go

// WRITING ActivityStreams
writer.Announce(
    writer.Actor().Name("John Doe").ID("john@doe.com").URL("john.doe.com"),
    writer.Document().URL("https://me.website.com/document-name"),
    nil
)

// READING ActivityStreams
object, err := New(`{
    "@context": "https://www.w3.org/ns/activitystreams",
    "name": "Foo",
    "id": "http://example.org/foo"
    }`)

object.Name() // "foo"
object.ID() // "http://example.org/foo"

```

## Pull Requests Welcome

This library is a work in progress, and will benefit from your experience reports, use cases, and contributions.  If you have an idea for making this library better, send in a pull request.  We're all in this together! 🏞
