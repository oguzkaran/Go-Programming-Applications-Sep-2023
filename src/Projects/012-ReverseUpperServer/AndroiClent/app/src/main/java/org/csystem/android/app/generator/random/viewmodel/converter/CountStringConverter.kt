package org.csystem.android.app.generator.random.viewmodel.converter

import androidx.databinding.InverseMethod

object CountStringConverter {
    @InverseMethod("toStr")
    fun toCount(str: String) : Int
    {
        var result = 0

        try {
            result = str.toUInt().toInt()
        }
        catch (_: NumberFormatException) {

        }

        return result
    }

    fun toStr(value: Int) = value.toString()
}