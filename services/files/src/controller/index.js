const { Video, Thubmnail, Info } = require('../db/models')
const { uploadInfo } = require('../db/query.js/upload')
// const fs = require('fs')
// const { Readable, Duplex } = require('stream');

const postFile = async (req, res, next) => {
    const {
        title,
        description,
        category,
        tags,
        uploadedUserId
    } = req.body
    try {
        const payload = {
            title,
            description,
            category,
            tags,
            uploadedUserId: Number(uploadedUserId),
            files: {
                thubmnailId: req.files.thubId,
                videoId: req.files.videoId
            }
        }
        const message = await uploadInfo(payload)
        res.status(200).end("video")
    } catch (err) {
        console.log(err)
    }
}

const stream = async (req, res) => {
    const { id } = req.params
    try {
        const data = await Video.find({ originalname: id })
        const arr = data.map(item => item.buffer)
        const buf = Buffer.concat(arr)
        res.set({
            'Content-Type': 'video/mp4',
            'Content-Length': buf.length,
            'Accept-Ranges': 'bytes',
            'Content-Range': `bytes ${0}-${buf.length}`,
        })
        res.end(buf)
    } catch (err) {
        console.log(err)
    }
}

const getFileImgaeByName = async (req, res, next) => {
    const { id } = req.params
    try {
        const data = await Thubmnail.find({ originalname: id })
        res.status(200).end(data[0].buffer)
    } catch (err) {
        console.log(err)
    }
}

const getFiles = async (req, res, next) => {
    try {
        const files = await Info.find({})
        res.status(200).json(files)
    } catch (err) {
        next(err)
    }
}

module.exports = {
    postFile,
    stream,
    getFiles,
    getFileImgaeByName
}