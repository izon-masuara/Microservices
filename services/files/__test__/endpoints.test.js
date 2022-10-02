const request = require('supertest')
const app = require('../app')
const connect = require('./db_test/config_test')
const mongoose = require('mongoose')
const path = require('path')
const { Info } = require('../src/db/models')

const baseUrl = `/api/v1/files`

describe("Test the root path", () => {
  beforeAll(() => {
    connect()
  })

  afterAll((done) => {
    mongoose.connection.db.dropDatabase((err, res) => {
      if (err) {
        console.log(err)
      } else {
        console.log(res)
      }
      done()
    })
  })

  it(`POST information without payload`, async () => {
    const response = await request(app)
      .post(`${baseUrl}/upload`)
    const { statusCode, body } = response
    expect(statusCode).toEqual(400)
    expect(Array.isArray(body.messageError)).toEqual(true)
  })

  it(`POST information without files`, async () => {
    const response = await request(app)
      .post(`${baseUrl}/upload`)
      .field('title', 'Video tutorial')
      .field('description', 'Ini merupakan video tutorial pembuatan kamus mengguankan node js')
      .field('category', 'Pembelajaran')
      .field('tags[0]', 'nodejs')
      .field('tags[1]', 'coding')
      .field('uploadedUserId', '20')
    const { statusCode, body } = response
    expect(statusCode).toEqual(400)
    expect(body.messageError).toBe('Must be upload thubmnail and video')
  })

  it(`POST information without one of files payload`, async () => {
    const picture = path.resolve(__dirname, "index.jpeg")
    const response = await request(app)
      .post(`${baseUrl}/upload`)
      .field('title', 'Video tutorial')
      .field('description', 'Ini merupakan video tutorial pembuatan kamus mengguankan node js')
      .field('category', 'Pembelajaran')
      .field('tags[0]', 'nodejs')
      .field('tags[1]', 'coding')
      .field('uploadedUserId', '20')
      .attach("files", picture)
    const { statusCode, body } = response
    expect(statusCode).toEqual(400)
    expect(body.messageError).toBe('Must be upload thubmnail and video')
  })

  it(`POST success upload information`, async () => {
    const picture = path.resolve(__dirname, "index.jpeg")
    const video = path.resolve(__dirname, "video.mp4")
    const response = await request(app)
      .post(`${baseUrl}/upload`)
      .field('title', 'Video tutorial')
      .field('description', 'Ini merupakan video tutorial pembuatan kamus mengguankan node js')
      .field('category', 'Pembelajaran')
      .field('tags[0]', 'nodejs')
      .field('tags[1]', 'coding')
      .field('uploadedUserId', '20')
      .attach("files", picture)
      .attach("files", video)
    const { statusCode, body } = response
    expect(statusCode).toEqual(201)
    expect(body.message).toBe('Success upload information')
  })

  it(`GET files to contain array`, async () => {
    const response = await request(app)
      .get(`${baseUrl}/`)
    const { statusCode, body } = response
    expect(statusCode).toEqual(200);
    expect(Array.isArray(body)).toEqual(true)
  });

  it('GET thubmnail with invalid id', async () => {
    let id = "khasdue813-=jsasdjlk"
    const response = await request(app)
    .get(`${baseUrl}/image/${id}`)
    const { statusCode,body } = response
    expect(statusCode).toEqual(404)
    expect(body.messageError).toBe("Image not found")
  })

  it('GET Success get thumbnail with id', async () => {
    let id = null
    const data = await Info.find({})
    id = data[0].files.thubmnailId
    const response = await request(app)
    .get(`${baseUrl}/image/${id}`)
    const { statusCode,body } = response
    expect(statusCode).toEqual(200)
  })

  it('GET video with invalid id', async () => {
    let id = "khasdue813-=jsasdjlk"
    const response = await request(app)
    .get(`${baseUrl}/video/${id}`)
    const { statusCode,body } = response
    expect(statusCode).toEqual(404)
    expect(body.messageError).toBe("Video not found")
  })

  it('GET Success get video with id', async () => {
    let id = null
    const data = await Info.find({})
    id = data[0].files.thubmnailId
    const response = await request(app)
    .get(`${baseUrl}/image/${id}`)
    const { statusCode,body } = response
    expect(statusCode).toEqual(200)
  })

})