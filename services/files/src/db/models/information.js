const { Schema, default: mongoose } = require('mongoose')

const information = new Schema({
    title: {
        type: String,
        require : true
    },
    description: {
        type: String,
        require : true
    },
    category: {
        type: String,
        require : true
    },
    tags: {
        type: Array,
        require : true
    },
    uploadedUserId: {
        type: Number,
        require : true
    },
    files : {
        thubmnailId:{
            type : String,
            require : true
        },
        videoId: {
            type : String,
            require : true
        },
    },
    createdAt: {
        type: Date,
        default: Date.now
    },
    updatedAt: {
        type: Date,
        default: Date.now
    }
})

const Info = mongoose.model('informations', information)

module.exports = Info