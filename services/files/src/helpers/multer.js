const multer = require('multer')
const { uploadFile } = require('../db/query.js/upload')
const upload = multer().array('files')

const requestFile = (req, res, next) => {
    upload(req, res, async function (err) {
        try {
            if (err) {
                next(err)
            }
            const thubId = await uploadFile(req.files[0])
            const videoId = await uploadFile(req.files[1])
            req.files = {
                thubId,
                videoId
            }
            next()
        } catch (error) {
            next(error)
        }
    })
}

module.exports = requestFile