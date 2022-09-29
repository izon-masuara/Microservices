const mongoose = require('mongoose')

const connect = async () => {
    try {
        await mongoose.connect('mongodb://localhost:27017/files')
        console.log("connect to database")    
    } catch (err) {
        console.error(err);
    }
}

module.exports = connect