const { Video,Thubmnail,Info } = require('../models')

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
            return video.originalname
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
        return err
    }
}

const uploadInfo = async(payload) => {
    try {
        const uploaded = await Info.create(payload)
        return uploaded
    } catch (err) {
        return err
    }
}

module.exports = {
    uploadFile,
    uploadInfo
}