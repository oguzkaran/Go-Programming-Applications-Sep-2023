package org.csystem.android.app.generator.random.viewmodel.converter

import androidx.databinding.InverseMethod

object PortStringConverter {
    @InverseMethod("toStr")
    fun toPort(str: String) : Int
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