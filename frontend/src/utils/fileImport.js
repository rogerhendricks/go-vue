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
                Object.assign(dataResult, result);
                resolve(dataResult);
            });
            reader.addEventListener('error', () => {
                reject('File read error');
            });

        } else if (file.name.endsWith(".log")) {
            reader.readAsText(file);
            reader.addEventListener('load', (e) => {
                const data = e.target.result;
                const result = csvToJSON(data);
                Object.assign(dataResult, result);
                resolve(dataResult);
            });
            reader.addEventListener('error', (e) => {
                reject(e);
            });
        } else if (file.name.endsWith(".bnk")) {
            reader.readAsText(file);
            reader.addEventListener('load', (e) => {
                const data = e.target.result;
                const result = bnkToJSON(data);
                Object.assign(dataResult, result);
                resolve(dataResult);
            });
            reader.addEventListener('error', (e) => {
                reject(e);
            });
        } else {
            reject('Unsupported file format');
        }
    });
}

//////////////////////////////////////// XML to JSON Conversion //////////////////////////////////////////
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
    // console.log(result);

    // Populate the global dataResult object
    return updateXmlDataResult(result);
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
    return dataResult;
}

// Format Date Helper
function formatDate(dtm) {
    return `${dtm.substring(0, 4)}-${dtm.substring(4, 6)}-${dtm.substring(6, 8)}`;
}
////////////////////////////////////////// CSV to JSON Conversion //////////////////////////////////////////
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
    
    return updateLogDataResult(result);
}

function updateLogDataResult(parsedData) {
    if (parsedData["102"] && parsedData["102"].value) {
        dataResult.serial = parsedData["102"].value;
    }
    // Update other relevant fields in dataResult
    // console.log(dataResult);
    return dataResult;
}
////////////////////////////////////////// BNK to JSON Conversion //////////////////////////////////////////
// BNK to JSON Conversion for .bnk files
function bnkToJSON(data) {
    let vals = data.split('\n'); // Split the CSV content into rows
    let result = {}; // Initialize result as an object to store key-value pairs

    for (let i in vals) {
        let row = vals[i].split(',');

        // Skip lines that don't have the expected key-value format
        if (row.length === 2) {
            let key = row[0].trim();   // Trim whitespace from the key
            let value = row[1].trim(); // Trim whitespace from the value

            // Store each key-value pair in the result object
            result[key] = value;
        }
    }

    console.log('result', result); // Output the parsed data for debugging

    // Now update the dataResult object with the parsed data
   return updateBnkDataResult(result);
}

// Function to update the dataResult object with the parsed CSV data
function updateBnkDataResult(parsedData, deviceSerial) {
    // Here are some examples of how to map parsed data to dataResult
    console.log('parsed data',parsedData); // Output the parsed data for debugging
    if (parsedData["SystemSerialNumber"] != deviceSerial) {
         alert("The serial number in the file does not match the device serial number.");
    }
    if (parsedData["SystemName"]) {
        dataResult.mdc_idc_set_brady_max_tracking_rate = parsedData["SystemName"];
    }
    if (parsedData["BdyPSBradyMode"]) {
        dataResult.mdc_idc_set_brady_mode = parsedData["BdyPSBradyMode"];
    }
    if (parsedData["NormParams.LRLIntvl"]) {
        let lrlToBpm =  Math.floor(60000 / parsedData["NormParams.LRLIntvl"].slice(0, -3));
        dataResult.mdc_idc_set_brady_lowrate= lrlToBpm;
    }
    if (parsedData["NormParams.MTRIntvl"]) {
        dataResult.mdc_idc_set_brady_max_tracking_rate = Math.floor(parsedData["NormParams.MTRIntvl"].slice(0, -3));
    }
    if (parsedData["NormParams.MSRIntvl"]) {
        dataResult.mdc_idc_set_brady_max_sensor_rate = Math.floor(parsedData["NormParams.MSRIntvl"].slice(0, -3));
    }
    if (parsedData["BatteryStatus.BatteryPhase"]) {
        dataResult.mdc_idc_batt_status = parsedData["BatteryStatus.BatteryPhase"];
    }

    // Update other relevant fields in dataResult based on the parsed data
     // Output the updated dataResult for debugging
     return dataResult;
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
