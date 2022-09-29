const { Schema, default: mongoose } = require('mongoose')

const videos = new Schema({
    originalname : {
        type: String,
        required: true
    },
    content_type : {
        type : String,
        required: true
    },
    queue : {
        type : Number,
        required : true
    },
    buffer : {
        type : Buffer,
        required: true
    },
    createdAt : {
        type : Date,
        default : Date.now
    },
})

const Video = mongoose.model('videos',videos)

module.exports = Video