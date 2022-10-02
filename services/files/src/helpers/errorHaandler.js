const errorHandler = async (err, req, res, next) => {
    switch (err.message.reason) {
        case "invalid key payload":
            res.status(err.code).json(err.message)
            break;
        case "file required":
            res.status(err.code).json(err.message)
            break;
        case "image extention":
            res.status(err.code).json(err.message)
            break;
        case "image size":
            res.status(err.code).json(err.message)
            break;
        case "video extention":
            res.status(err.code).json(err.message)
            break;
        case "not found":
            res.status(err.code).json(err.message)
            break;
        default:
            res.status(500).json({ message: "Internal server error" })
            break;
    }
}

module.exports = errorHandler