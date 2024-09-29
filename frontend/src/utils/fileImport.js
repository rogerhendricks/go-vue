export default function fileImport(file, deviceSerial) {
    return new Promise((resolve, reject) => {
        console.log(file.name, file.type, deviceSerial);
        const reader = new FileReader();

        if (file.type === 'text/xml') {
            reader.readAsText(file);
            reader.addEventListener('load', () => {
                const data = reader.result;
                const parser = new DOMParser();
                const xmlDoc = parser.parseFromString(data, 'text/xml');
                const result = xmlToJson(xmlDoc);
                resolve(result);
            });
            reader.addEventListener('error', () => {
                reject('File read error');
            });

        } else if (file.name.endsWith(".log")) {
            reader.readAsText(file);
            reader.addEventListener('load', (e) => {
                const data = e.target.result;
                const result = csvToJSON(data);
                resolve(result); // Ensure to resolve after processing
            });
            reader.addEventListener('error', (e) => {
                reject(e);
            });

        } else {
            reject('Unsupported file format');
        }
    });
}

// XML to JSON Conversion
function xmlToJson(xmlDoc) {
    let result = { };

    let dev = evaluateXPath(xmlDoc, 'biotronik-ieee11073-export/dataset/section[@name="MDC"]/section[@name="IDC"]/section[@name="DEV"]//value');
    let sess = evaluateXPath(xmlDoc, 'biotronik-ieee11073-export/dataset/section[@name="MDC"]/section[@name="IDC"]/section[@name="SESS"]//value');
    let msmtBatt = evaluateXPath(xmlDoc, 'biotronik-ieee11073-export/dataset/section[@name="MDC"]/section[@name="IDC"]/section[@name="MSMT"]//section[@name="BATTERY"]//value');
    let msmtRa = evaluateXPath(xmlDoc, 'biotronik-ieee11073-export/dataset/section[@name="MDC"]/section[@name="IDC"]/section[@name="MSMT"]//section[@name="LEADCHNL_RA"]//value');
    let msmtRv = evaluateXPath(xmlDoc, 'biotronik-ieee11073-export/dataset/section[@name="MDC"]/section[@name="IDC"]/section[@name="MSMT"]//section[@name="LEADCHNL_RV"]//value');
    let brady = evaluateXPath(xmlDoc, 'biotronik-ieee11073-export/dataset/section[@name="MDC"]/section[@name="IDC"]/section[@name="SET"]//section[@name="BRADY"]//value');


    // Parsing XML into respective objects
    // result['dev'] = parseNode(dev);
    // result['sess'] = parseNode(sess);
    // result['msmtBatt'] = parseNode(msmtBatt);
    // result['msmtRa'] = parseNode(msmtRa);
    // result['msmtRv'] = parseNode(msmtRv);
    // result['brady'] = parseNode(brady);
    result['dev'] = dev;
    result['sess'] = sess;
    result['msmtBatt'] = msmtBatt;
    result['msmtRa'] = msmtRa;
    result['msmtRv'] = msmtRv;
    result['brady'] = brady;
    console.log(result);

    // Populate the global dataResult object
    updateXmlDataResult(result);

    return result;
}

// Helper function to extract XPath results into an object
function evaluateXPath(xmlDoc, xpath) {
    let result = {};
    let nodes = xmlDoc.evaluate(xpath, xmlDoc, null, XPathResult.ANY_TYPE, null);
    let node;
    while (node = nodes.iterateNext()) {
        result[node.getAttribute('name')] = node.textContent;
    }
    return result;
}

// Populate dataResult object based on parsed XML
function updateXmlDataResult(parsedData) {
    dataResult.serial = parsedData.dev.SERIAL;
    dataResult.report_date = formatDate(parsedData.sess.DTM);

    // Update dataResult with the XML data
    dataResult.mdc_idc_dev_model = parsedData.dev.MODEL;
    dataResult.mdc_idc_dev_sav = parsedData.dev.SAV;
    dataResult.mdc_idc_dev_pav = parsedData.dev.PAV;
    dataResult.mdc_idc_set_brady_mode = parsedData.brady.MODE;
    dataResult.mdc_idc_batt_volt = parsedData.msmtBatt.VOLTAGE;
    console.log(dataResult);    
}

// Format Date Helper
function formatDate(dtm) {
    return `${dtm.substring(0, 4)}-${dtm.substring(4, 6)}-${dtm.substring(6, 8)}`;
}

// function parseNode(node) {
//     let obj = {};

//     // Check if the node is an element node (type 1)
//     if (node.nodeType === 1) {
//         // Parse attributes
//         if (node.attributes && node.attributes.length > 0) {
//             for (let i = 0; i < node.attributes.length; i++) {
//                 const attribute = node.attributes.item(i);
//                 obj[attribute.nodeName] = attribute.nodeValue;
//             }
//         }

//         // Parse child nodes if it has children
//         if (node.hasChildNodes()) {
//             for (let j = 0; j < node.childNodes.length; j++) {
//                 const child = node.childNodes.item(j);
//                 const childName = child.nodeName;

//                 // If the object doesn't already have this child name, create it
//                 if (typeof obj[childName] === "undefined") {
//                     obj[childName] = parseNode(child);
//                 } else {
//                     // If it already exists, make it an array if it's not already one
//                     if (!Array.isArray(obj[childName])) {
//                         obj[childName] = [obj[childName]];
//                     }
//                     obj[childName].push(parseNode(child));
//                 }
//             }
//         } else {
//             obj = node.textContent;
//         }
//     }
//     return obj;
// }


// CSV to JSON Conversion for .log files
function csvToJSON(data) {
    const result = {}; 
    const FS = String.fromCharCode(28); // File separator; end of file. Or between a concatenation of what might otherwise be separate files.
    const GS = String.fromCharCode(29); // Group separator; between sections of data. Not needed in simple data files.
    const RS = String.fromCharCode(30); // Record separator; end of a record or row.
    const US = String.fromCharCode(31); // Unit separator; between fields of a record, or members of a row.
    let vals = data.split("\n");
    for (let i = 0; i < vals.length; i++) {
        let row = vals[i].split(FS);
        let code = row[0].toString();
        result[code] = { name: row[1], value: row[2] };
    }

    // Assuming the CSV data updates specific parts of dataResult
    
    updateLogDataResult(result);
    return result;
}

function updateLogDataResult(parsedData) {
    if (parsedData["102"] && parsedData["102"].value) {
        dataResult.serial = parsedData["102"].value;
    }
    // Update other relevant fields in dataResult
    console.log(dataResult);
}
// Initial dataResult template
const dataResult = {
    report_date: "",
    mdc_idc_stat_ataf_burden_percent: "",
    mdc_idc_stat_pvc_count: "",
    mdc_idc_stat_nsvt_count: "",
    mdc_idc_set_brady_mode: "",
    mdc_idc_set_brady_lowrate: "",
    mdc_idc_set_brady_max_tracking_rate: "",
    mdc_idc_set_brady_max_sensor_rate: "",
    mdc_idc_dev_sav: "",
    mdc_idc_dev_pav: "",
    mdc_idc_stat_brady_ra_percent_paced: "",
    mdc_idc_stat_brady_rv_percent_paced: "",
    mdc_idc_stat_brady_lv_percent_paced: "",
    biv_percent_paced: "",
    mdc_idc_batt_volt: "",
    mdc_idc_batt_remaining: "",
    mdc_idc_batt_status: "",
    mdc_idc_cap_charge_time: "",
    mdc_idc_msmt_ra_impedance_mean: "",
    mdc_idc_msmt_ra_sensing_mean: "",
    mdc_idc_msmt_ra_threshold: "",
    mdc_idc_msmt_ra_pw: "",
    mdc_idc_msmt_rv_impedance_mean: "",
    mdc_idc_msmt_rv_sensing_mean: "",
    mdc_idc_msmt_rv_threshold: "",
    mdc_idc_msmt_rv_pw: "",
    mdc_idc_msmt_lv_impedance_mean: "",
    mdc_idc_msmt_lv_sensing_mean: "",
    mdc_idc_msmt_lv_threshold: "",
    mdc_idc_msmt_lv_pw: "",
};
