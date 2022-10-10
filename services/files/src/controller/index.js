const { response } = require('../../app')
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
                videoId: req.files.videoId,
                size : req.files.size
            }
        }
        await uploadInfo(payload)
        res.status(201).json({
            message: "Success upload information"
        })
    } catch (err) {
        next(err)
    }
}

const stream = async (req, res, next) => {
    const { id } = req.params
    try {
        const data = await Video.find({ originalname: id })
        const arr = data.map(item => item.buffer)
        const buf = Buffer.concat(arr)
        if (data.length == 0) {
            throw {
                code: 404,
                message: {
                    reason: "not found",
                    messageError: "Video not found"
                }
            }
        }
        res.set({
            'Content-Type': 'video/mp4',
            'Content-Length': buf.length,
            'Accept-Ranges': 'bytes',
            'Content-Range': `bytes ${0}-${buf.length}`,
        })
        res.status(200).end(buf)
    } catch (err) {
        next(err)
    }
}

const getFileImgaeByName = async (req, res, next) => {
    const { id } = req.params
    try {
        const data = await Thubmnail.find({ originalname: id })
        res.status(200).end(data[0].buffer)
    } catch (err) {
        const msgErr = {
            code: 404,
            message: {
                reason: "not found",
                messageError: "Image not found"
            }
        }
        next(msgErr)
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