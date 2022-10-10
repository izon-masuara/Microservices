const { Video,Thubmnail,Info } = require('../models')

/**
 * 
 * @param {Object} payload Object of file from multer 
 * 
 * @return {string} Original name
 * 
 * This fuction will upload file base on mimetype
 */
const uploadFile = async (payload) => {
    const {
        originalname,
        mimetype,
        buffer
    } = payload
    try {
        if (mimetype == "video/mp4") {
            size = buffer.length
            start = 0
            end = 15000000
            totalSize = []
            while(size){
                if(size > 15000000){
                    size = size - 15000000
                    totalSize.push(buffer.slice(start,end))
                    start = end
                    end = end + 15000000
                }else {
                    totalSize.push(buffer.slice(start,end))
                    break
                }
            }
            const video = {
                originalname : `${Date.now()}-wppq-video-${Math.random() * 10}-${originalname}`,
                content_type : mimetype,
            }
            for(i = 0; i < totalSize.length; i++){
                video.buffer = totalSize[i]
                video.queue = 1 + i
                await Video.create(video)
            }
            return {
                videoName : video.originalname,
                size : buffer.length
            }
        }else if (mimetype == "image/jpeg" || mimetype == "image/jpg" || mimetype == "image/png"){
            const thubmnail = {
                originalname : `${Date.now()}-wppq-thumbnail-${Math.random() * 10}-${originalname}`,
                content_type : mimetype,
                buffer
            }
            const data = await Thubmnail.create(thubmnail)
            return data.originalname
        }
    } catch (err) {
        throw {
            code : 400,
            message : err
        }
    }
}

/**
 * 
 * @param {Object} payload Schema of Information model
 * @returns {Object} message
 */
const uploadInfo = async(payload) => {
    try {
        const uploaded = await Info.create(payload)
        return uploaded
    } catch (err) {
        throw {
            code : 400,
            message : err
        }
    }
}

module.exports = {
    uploadFile,
    uploadInfo
}