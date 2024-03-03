package org.csystem.android.app.generator.random.viewmodel.converter

import androidx.databinding.InverseMethod

object MaximumStringConverter {
    @InverseMethod("toStr")
    fun toMaximum(str: String) : Int
    {
        var result = 0

        try {
            result = str.toInt();
        }
        catch (_: NumberFormatException) {

        }

        return result
    }

    fun toStr(value: Int) = value.toString()
}