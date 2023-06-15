# shitchain

waow much cosmos very learning such development

# Overview

* To start from scratch can be scaffold with `ignite scaffold chain github.com/bonedaddy/shitchain`
* To scaffold a webui run `ignite scaffold vue`
* To scaffold a new module that is IBC enabled `ignite scaffold module shitstore --ibc`
* To scaffold a list type named `Post` run `ignite scaffold list post title body amount:coin tags:array.string --module shitstore`
* To scaffold a new message type called `BroadcastPost` run `ignite scaffold message broadcast-post post:Post --module shitstore --response id:int,title:string`
* To scaffold a new query called `CountPosts` run `ignite scaffold query count-posts creator --response number:int --module shitstore --paginated`
* To send a transaction for the `CreatePost` message, run `shitchaind tx shitstore create-post "this is a test post" "this is a test body" 100stake "foo,bar,baz" --from alice`
* To send a transaction for the `BroadcastPost` message run `shitchaind tx shitstore broadcast-post '{"title": "foo", "body": "bar", "amount": {}, "tags": []}'  --from alice `
* To list all known posts run `shitchaind query shitstore list-post`
* To run our custom query `shitchaind query shitstore count-posts cosmos10k9kaamtqf03law4cm9rntwq9v6hk8uzkh7s8l`
# Resources

* https://docs.ignite.com/guide
* https://docs.ignite.com/