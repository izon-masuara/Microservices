const { Schema, default: mongoose } = require('mongoose')

const thubmnail = new Schema({
    originalname : {
        type: String,
        required: true
    },
    content_type : {
        type : String,
        required: true
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

const Thubmnail = mongoose.model('thubmnails',thubmnail)

module.exports = Thubmnail