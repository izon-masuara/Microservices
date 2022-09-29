const multer = require('multer')
const { uploadFile } = require('../db/query.js/upload')
const upload = multer().array('files')

const requestFile = (req, res, next) => {
    upload(req, res, async function (err) {
        try {
            if (err) {
                throw err
            }

            //  Validation key of payload
            messageError = []
            const keys = ["title","description","category","tags","uploadedUserId"]

            keys.forEach(key => {
                if(key in req.body == false){
                    messageError.push(key + " required")
                }
            })

            if (messageError.lenght == 0) {
                throw {
                    code : 400,
                    message : {
                        reason : "invalid key payload",
                        messageError
                    }
                }
            }

            const thubmnail = req.files[0]
            const video = req.files[1]

            //  Vilidation files
            if(thubmnail == undefined || video == undefined) {
                throw {
                    code : 400,
                    message : {
                        reason : "file required",
                        msg : "Must be upload thubmnail and video"
                    }
                }
            }

            if(!thubmnail.mimetype == "image/jpeg" || !thubmnail.mimetype == "image/jpg" || !thubmnail.mimetype == "image/png"){
                throw {
                    code : 400,
                    message : {
                        reason : "image extention",
                        msg : "Image file must be .jpeg or .jpg or .png"
                    }
                }
            }else if (thubmnail.size >= 15000000) {
                throw {
                    code : 400,
                    message : {
                        reason : "image size",
                        msg : "image size must be less than 15Mb"
                    }
                }
            }

            if(video.mimetype != "video/mp4"){
                throw {
                    code : 400,
                    message : {
                        reason : "video extention",
                        msg : "Vide file must be .mp4"
                    }
                }
            }

            // Upload video and thubmnail
            const thubId = await uploadFile(thubmnail)
            const videoId = await uploadFile(video)

            // Data will pass to the next middleware
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