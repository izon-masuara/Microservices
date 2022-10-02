const mongoose = require('mongoose')

const connect = async () => {
    try {
        await mongoose.connect(`${process.env.MONGODB_URI}/files`)
        console.log("connect to database")    
    } catch (err) {
        console.error(err);
    }
}

module.exports = connect