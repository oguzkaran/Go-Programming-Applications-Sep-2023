package org.csystem.android.app.generator.random.viewmodel.converter

import androidx.databinding.InverseMethod

object MinimumStringConverter {
    @InverseMethod("toStr")
    fun toMinimum(str: String) : Int
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