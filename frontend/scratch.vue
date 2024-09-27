const handleSubmit = async () => {
  try {
    const uploadData = new FormData()
    uploadData.append('patient_id', formData.value.patient)
    uploadData.append('report_date', formData.value.report_date)

    // Check if new files are selected
    if (formData.value.files && formData.value.files.length > 0) {
      // Create a new PDF document to merge files
      const mergedPdfDoc = await PDFDocument.create()

      // Loop through each selected file and merge them
      for (const file of formData.value.files) {
        const arrayBuffer = await file.arrayBuffer()
        const pdfDoc = await PDFDocument.load(arrayBuffer)
        const copiedPages = await mergedPdfDoc.copyPages(pdfDoc, pdfDoc.getPageIndices())
        copiedPages.forEach((page) => mergedPdfDoc.addPage(page))
      }

      // Serialize the merged PDF into a Blob
      const mergedPdfBytes = await mergedPdfDoc.save()
      const mergedPdfBlob = new Blob([mergedPdfBytes], { type: 'application/pdf' })

      // Append the merged PDF Blob to FormData
      uploadData.append('file', mergedPdfBlob, 'merged_report.pdf')
    }

    // Send the FormData with or without the merged PDF to the backend
    const response = await axios.put(`/api/reports/${props.report_id}/`, uploadData, {
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    })

    // Update the report in the Vuex store
    store.commit('setReport', response.data.report)
    console.log('Report updated:', response.data.report)
    toast.success('Report updated successfully', { timeout: 2000 })
  } catch (error) {
    console.error('Error updating report:', error)
    toast.error('Error updating report', { timeout: 2000 })
  }
}