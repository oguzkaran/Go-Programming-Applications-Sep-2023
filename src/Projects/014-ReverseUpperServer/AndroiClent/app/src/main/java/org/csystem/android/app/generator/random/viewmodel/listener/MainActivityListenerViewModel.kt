package org.csystem.android.app.generator.random.viewmodel.listener

import org.csystem.android.app.generator.random.MainActivity
import java.lang.ref.WeakReference

class MainActivityListenerViewModel(activity: MainActivity) {
    private val mWeakReference = WeakReference(activity)

    fun handleGenerateButton() = mWeakReference.get()?.generateButtonClicked()
    fun handleConfigurationButton() = mWeakReference.get()?.configurationButtonClicked()
    fun handleSaveButton() = mWeakReference.get()?.saveButtonClicked()
    fun handleClearButton() = mWeakReference.get()?.clearButtonClicked()
}