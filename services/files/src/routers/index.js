const router = require('express').Router()
const { postFile,
        stream,
        getFiles,
        getFileImgaeByName
} = require('../controller')
const errorHandler = require('../helpers/errorHaandler')
const requestFile = require('../helpers/multer')

const baseUrl = `/api/v1/files`

router.post(`${baseUrl}/upload`, requestFile, postFile)
router.get(`${baseUrl}/`, getFiles)
router.get(`${baseUrl}/image/:id`,getFileImgaeByName)
router.get(`${baseUrl}/video/:id`, stream)

router.use(errorHandler)

module.exports = router