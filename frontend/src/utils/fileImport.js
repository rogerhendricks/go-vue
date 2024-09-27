export default function fileImport(file, deviceSerial) {
    return new Promise((resolve, reject) => {
        console.log(file.name, file.type, deviceSerial)
        return;
    })
}